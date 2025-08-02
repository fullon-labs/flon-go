package snapshot

import "io"

type SectionName string

const (
	SectionNameChainSnapshotHeader         SectionName = "flon::chain::chain_snapshot_header"
	SectionNameBlockState                  SectionName = "flon::chain::block_state"
	SectionNameAccountObject               SectionName = "flon::chain::account_object"
	SectionNameAccountMetadataObject       SectionName = "flon::chain::account_metadata_object"
	SectionNameAccountRamCorrectionObject  SectionName = "flon::chain::account_ram_correction_object"
	SectionNameGlobalPropertyObject        SectionName = "flon::chain::global_property_object"
	SectionNameProtocolStateObject         SectionName = "flon::chain::protocol_state_object"
	SectionNameDynamicGlobalPropertyObject SectionName = "flon::chain::dynamic_global_property_object"
	SectionNameBlockSummaryObject          SectionName = "flon::chain::block_summary_object"
	SectionNameTransactionObject           SectionName = "flon::chain::transaction_object"
	SectionNameGeneratedTransactionObject  SectionName = "flon::chain::generated_transaction_object"
	SectionNameCodeObject                  SectionName = "flon::chain::code_object"
	SectionNameContractTables              SectionName = "contract_tables"
	SectionNamePermissionObject            SectionName = "flon::chain::permission_object"
	SectionNamePermissionLinkObject        SectionName = "flon::chain::permission_link_object"
	SectionNameResourceLimitsObject        SectionName = "flon::chain::resource_limits::resource_limits_object"
	SectionNameResourceUsageObject         SectionName = "flon::chain::resource_limits::resource_usage_object"
	SectionNameResourceLimitsStateObject   SectionName = "flon::chain::resource_limits::resource_limits_state_object"
	SectionNameResourceLimitsConfigObject  SectionName = "flon::chain::resource_limits::resource_limits_config_object"
	SectionNameGenesisState                SectionName = "flon::chain::genesis_state"

	// Ultra Specific
	SectionAccountFreeActionsObject SectionName = "flon::chain::account_free_actions_object"
)

type Section struct {
	Name       SectionName
	Offset     uint64
	Size       uint64 // This includes the section name and row count
	BufferSize uint64 // This represents the bytes that are following the section header
	RowCount   uint64 // This is a count of rows packed in `Buffer`
	Buffer     io.Reader
}

type sectionHandlerFunc func(s *Section, f sectionCallbackFunc) error
type sectionCallbackFunc func(obj interface{}) error
