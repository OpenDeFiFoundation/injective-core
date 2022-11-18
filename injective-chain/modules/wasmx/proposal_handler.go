package wasmx

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/InjectiveLabs/injective-core/injective-chain/modules/wasmx/keeper"
	"github.com/InjectiveLabs/injective-core/injective-chain/modules/wasmx/types"
)

// NewWasmxProposalHandler creates a governance handler to manage new wasmx proposal types.
func NewWasmxProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ContractRegistrationRequestProposal:
			return handleContractRegistrationRequestProposal(ctx, k, c)
		case *types.BatchContractRegistrationRequestProposal:
			return handleBatchContractRegistrationRequestProposal(ctx, k, c)
		case *types.BatchContractDeregistrationProposal:
			return handleBatchContractDeregistrationProposal(ctx, k, c)
		case *types.BatchStoreCodeProposal:
			return handleBatchStoreCodeProposal(ctx, k, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized wasmx proposal content type: %T", c)
		}
	}
}

func handleContractRegistration(ctx sdk.Context, k keeper.Keeper, params types.Params, req types.ContractRegistrationRequest) error {
	contractAddress, _ := sdk.AccAddressFromBech32(req.ContractAddress)

	// Enforce MinGasContractExecution ≤ GasLimit ≤ MaxContractGasLimit
	if req.GasLimit < types.MinExecutionGasLimit || req.GasLimit > params.MaxContractGasLimit {
		return sdkerrors.Wrapf(types.ErrInvalidGasLimit, "ContractRegistrationRequestProposal: The gasLimit (%d) must be within the range (%d) - (%d).", req.GasLimit, types.MinExecutionGasLimit, params.MaxContractGasLimit)
	}

	// Enforce GasPrice ≥ MinGasPrice
	if req.GasPrice < params.MinGasPrice {
		return sdkerrors.Wrapf(types.ErrInvalidGasPrice, "ContractRegistrationRequestProposal: The gasPrice (%d) must be greater than (%d)", req.GasPrice, params.MinGasPrice)
	}

	// Enforce that the registry contract exists
	registryContract, err := sdk.AccAddressFromBech32(params.RegistryContract)
	if err != nil || !k.DoesContractExist(ctx, registryContract) {
		return sdkerrors.Wrap(types.ErrInvalidContractAddress, "ContractRegistrationRequestProposal: The registry contract address is not set in params.")
	}

	if contractAddress.Equals(registryContract) {
		return sdkerrors.Wrapf(types.ErrInvalidContractAddress, "ContractRegistrationRequestProposal: The contract address was the same as registry contract address: %s", contractAddress.String())
	}

	// Enforce that a contract exists at contractAddress
	if !k.DoesContractExist(ctx, contractAddress) {
		return sdkerrors.Wrapf(types.ErrInvalidContractAddress, "ContractRegistrationRequestProposal: The contract address %s does not exist", contractAddress.String())
	}

	// Obtain the list of RawContractExecutionParams for registered contracts
	contractExecutionList, err := k.FetchRegisteredContractExecutionList(ctx, registryContract, false)
	if err != nil {
		return sdkerrors.Wrapf(err, "ContractRegistrationRequestProposal: Error fetching registered contracts\n")
	}

	// Enforce that the contract is not already registered
	for _, registeredContract := range contractExecutionList {
		registeredContractAddr, _ := sdk.AccAddressFromBech32(registeredContract.Address)
		if contractAddress.Equals(registeredContractAddr) {
			return sdkerrors.Wrapf(types.ErrAlreadyRegistered, "ContractRegistrationRequestProposal: contract %s is already registered", registeredContract.Address)
		}
	}

	// Register the contract execution parameters
	if err = k.RegisterContract(ctx, registryContract, req); err != nil {
		return sdkerrors.Wrapf(err, "ContractRegistrationRequestProposal: Error while registering the contract")
	}

	// Pin the contract with Wasmd module to reduce the gas used for contract execution
	if req.PinContract {
		if err = k.PinContract(ctx, contractAddress); err != nil {
			return sdkerrors.Wrapf(err, "ContractRegistrationRequestProposal: Error while pinning the contract")
		}
	}

	return nil
}

func handleContractRegistrationRequestProposal(ctx sdk.Context, k keeper.Keeper, p *types.ContractRegistrationRequestProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	params := k.GetParams(ctx)
	return handleContractRegistration(ctx, k, params, p.ContractRegistrationRequest)
}

func handleBatchContractRegistrationRequestProposal(ctx sdk.Context, k keeper.Keeper, p *types.BatchContractRegistrationRequestProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	params := k.GetParams(ctx)

	for _, req := range p.ContractRegistrationRequests {
		if err := handleContractRegistration(ctx, k, params, req); err != nil {
			return err
		}
	}

	return nil
}

func handleBatchContractDeregistrationProposal(ctx sdk.Context, k keeper.Keeper, p *types.BatchContractDeregistrationProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	params := k.GetParams(ctx)

	registryContract := sdk.MustAccAddressFromBech32(params.RegistryContract)
	contractExecutionList, err := k.FetchRegisteredContractExecutionList(ctx, registryContract, false)
	if err != nil {
		return sdkerrors.Wrapf(err, "BatchContractDeregistrationProposal: Error fetching registered contracts")
	}

	activeContracts := make(map[string]struct{})

	for _, c := range contractExecutionList {
		activeContracts[sdk.MustAccAddressFromBech32(c.Address).String()] = struct{}{}
	}

	for _, contract := range p.Contracts {
		contractAddress := sdk.MustAccAddressFromBech32(contract)

		// skip deregistration of contracts that don't exist [or if it's address is the same as registry contract's]
		if _, ok := activeContracts[contractAddress.String()]; !ok || contractAddress.Equals(registryContract) {
			continue
		}

		if err := k.DeregisterContract(ctx, registryContract, contractAddress); err != nil {
			return err
		}

		if err := k.UnpinContract(ctx, contractAddress); err != nil {
			return err
		}
	}

	return nil
}

func handleBatchStoreCodeProposal(ctx sdk.Context, k keeper.Keeper, p *types.BatchStoreCodeProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	proposalHandler := k.GetWasmProposalHandler()

	for _, req := range p.Proposals {
		if err := proposalHandler(ctx, &req); err != nil {
			return err
		}
	}

	return nil
}
