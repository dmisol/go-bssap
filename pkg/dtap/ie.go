package dtap

type DtapIE uint16

/*
	The same Type/Reference and Information Element can have different FORMATs and IEIs.
	Therefore, each tag is identified by name + format. For example, REJ_CAUSE_V.
*/

const (
	//general created Tags
	PROTOCOL_DISC_V                                 DtapIE = 0x100
	SKIP_IND_V                                      DtapIE = 0x101
	MSG_TYPE_V                                      DtapIE = 0x102
	AUTH_PARAM_RAND_V                               DtapIE = 0x105
	AUTH_RESP_PARAM_V                               DtapIE = 0x106
	REJ_CAUSE_V                                     DtapIE = 0x107
	PD_AND_SAPI_V                                   DtapIE = 0x10A
	CM_SERVICE_TYPE_V                               DtapIE = 0x10B
	IDENTITY_TYPE_V                                 DtapIE = 0x10C
	MOBILE_STATION_CLASSMARK_V                      DtapIE = 0x10D
	LOCATION_UPDATING_TYPE_V                        DtapIE = 0x10E
	LOCATION_AREA_ID_V                              DtapIE = 0x10F
	MS_NETWORK_FEATURE_SUPPORT_TV                   DtapIE = 0x110
	NON_3GPP_NW_PROVIDED_POLICIES_TV                DtapIE = 0x111
	P_TMSI_TYPE_TV                                  DtapIE = 0x112
	DEVICE_PROPS_TV                                 DtapIE = 0x113
	PAGE_MODE_V                                     DtapIE = 0x114
	DEDICATED_MODE_OR_TBF_V                         DtapIE = 0x115
	CHANNEL_DESC_V                                  DtapIE = 0x116
	PACKET_CHANNEL_DESC_V                           DtapIE = 0x117
	REQUEST_REF_V                                   DtapIE = 0x118
	TIMING_ADVANCE_V                                DtapIE = 0x119
	MOBILE_ALLOC_LV                                 DtapIE = 0x11A
	IA_REST_OCTETS_V                                DtapIE = 0x11B
	FEATURE_INDICATOR_V                             DtapIE = 0x11C
	REQUEST_REF_1_V                                 DtapIE = 0x11D
	WAIT_INDICATION_1_V                             DtapIE = 0x11E
	REQUEST_REF_2_V                                 DtapIE = 0x11F
	WAIT_INDICATION_2_V                             DtapIE = 0x120
	REQUEST_REF_3_V                                 DtapIE = 0x121
	WAIT_INDICATION_3_V                             DtapIE = 0x122
	REQUEST_REF_4_V                                 DtapIE = 0x123
	WAIT_INDICATION_4_V                             DtapIE = 0x124
	IAR_REST_OCTETS_V                               DtapIE = 0x125
	CIPHERING_MODE_SETTING_V                        DtapIE = 0x126
	CIPHER_RESPONSE_V                               DtapIE = 0x127
	RR_CAUSE_V                                      DtapIE = 0x128
	DESCRIPTION_OF_FIRST_CHANNEL_AFTER_TIME_V       DtapIE = 0x129
	POWER_COMMAND_V                                 DtapIE = 0x12A
	CIPHER_MODE_SETTING_TV                          DtapIE = 0x12B
	CELL_DESCRIPTION_V                              DtapIE = 0x12C
	HANDOVER_REFERENCE_V                            DtapIE = 0x12D
	POWER_COMMAND_AND_ACCESS_TYPE_V                 DtapIE = 0x12E
	SYNCHRONIZATION_INDICATION_TV                   DtapIE = 0x12F
	PLMN_INDEX_TV                                   DtapIE = 0x130
	PS_CAUSE_TV                                     DtapIE = 0x131
	GROUP_CIPHER_KEY_NUMBER_TV                      DtapIE = 0x132
	GPRS_RESUMPTION_TV                              DtapIE = 0x133
	ENHANCED_DTM_CS_RELEASE_INDICATION_TV           DtapIE = 0x134
	CHANNELS_NEEDED_FOR_MOBILES_1_AND_2_V           DtapIE = 0x135
	MOBILE_IDENTITY_1_LV                            DtapIE = 0x136
	P1_REST_OCTETS_V                                DtapIE = 0x137
	MOBILE_IDENTITY_1_V                             DtapIE = 0x138
	MOBILE_IDENTITY_2_V                             DtapIE = 0x139
	P2_REST_OCTETS_V                                DtapIE = 0x13A
	MOBILE_IDENTITY_3_V                             DtapIE = 0x13B
	MOBILE_IDENTITY_4_V                             DtapIE = 0x13C
	P3_REST_OCTETS_V                                DtapIE = 0x13D
	CIPHERING_KEY_SEQUENCE_NUMBER_V                 DtapIE = 0x13E
	SPARE_HALF_OCTET_V                              DtapIE = 0x13F
	MOBILE_STATION_CLASSMARK_LV                     DtapIE = 0x140
	MOBILE_IDENTITY_LV                              DtapIE = 0x141
	ADDITIONAL_UPDATE_PARAMETERS_TV                 DtapIE = 0x142
	TRANSACTION_IDEN_V                              DtapIE = 0x143
	BC_REPEAT_INDICATOR_TV                          DtapIE = 0x144
	LLC_REPEAT_INDICATOR_TV                         DtapIE = 0x145
	HLC_REPEAT_INDICATOR_TV                         DtapIE = 0x146
	CAUSE_LV                                        DtapIE = 0x147
	PRIORITY_TV                                     DtapIE = 0x148
	AUTH_PARAM_AUTN_TLV                             DtapIE = 0x20 //MM embedded tags
	AUTH_RESP_PARAM_EXT_TLV                         DtapIE = 0x21
	AUTH_FAILURE_PARAM_TLV                          DtapIE = 0x22
	LOCATION_AREA_ID_TV                             DtapIE = 0x13
	MM_TIMER_TLV                                    DtapIE = 0x36
	ROUTING_AREA_IDENTIFICATION_TLV                 DtapIE = 0x1B
	P_TMSI_SIGNATURE_TLV                            DtapIE = 0x19
	MOBILE_IDENTITY_TLV                             DtapIE = 0x17
	FOLLOW_ON_PROCEED_T                             DtapIE = 0xA1
	CTS_PERMISSION_T                                DtapIE = 0xA2
	EQUIVALENT_PLMNS_TLV                            DtapIE = 0x4A
	EMERGENCY_NUMBER_LIST_TLV                       DtapIE = 0x34
	PER_MS_T3212_TLV                                DtapIE = 0x35
	MOBILE_STATION_CLASSMARK_FOR_UMTS_TLV           DtapIE = 0x33
	FULL_NAME_FOR_NETWORK_TLV                       DtapIE = 0x43
	SHORT_NAME_FOR_NETWORK_TLV                      DtapIE = 0x45
	LOCAL_TIME_ZONE_TV                              DtapIE = 0x46
	UNIVERSAL_TIME_AND_LOCAL_TIME_ZONE_TV           DtapIE = 0x47
	LSA_IDENTITY_TLV                                DtapIE = 0x48
	NETWORK_DAYLIGHT_SAVING_TIME_TLV                DtapIE = 0x49
	STARTING_TIME_TV                                DtapIE = 0x7C //RR embedded tags
	EXTENDED_TSC_SET_TV                             DtapIE = 0x6D
	ME_IDENTITY_TLV                                 DtapIE = 0x17
	FREQUENCY_LIST_AFTER_TIME_TLV                   DtapIE = 0x05
	CELL_CHANNEL_DESCRIPTION_TV                     DtapIE = 0x62
	DESCRIPTION_OF_MULTISLOT_CONFIGURATION_TLV      DtapIE = 0x10
	MODE_OF_FIRST_CHANNEL_CHANNEL_SET_1_TV          DtapIE = 0x63
	MODE_OF_CHANNEL_SET_2_TV                        DtapIE = 0x11
	MODE_OF_CHANNEL_SET_3_TV                        DtapIE = 0x13
	MODE_OF_CHANNEL_SET_4_TV                        DtapIE = 0x14
	MODE_OF_CHANNEL_SET_5_TV                        DtapIE = 0x15
	MODE_OF_CHANNEL_SET_6_TV                        DtapIE = 0x16
	MODE_OF_CHANNEL_SET_8_TV                        DtapIE = 0x18
	DESCRIPTION_OF_SECOND_CHANNEL_AFTER_TIME_TV     DtapIE = 0x64
	MODE_OF_SECOND_CHANNEL_TV                       DtapIE = 0x66
	MOBILE_ALLOCATION_AFTER_TIME_TLV                DtapIE = 0x72
	FREQUENCY_LIST_BEFORE_TIME_TLV                  DtapIE = 0x19
	DESCRIPTION_OF_FIRST_CHANNEL_BEFORE_TIME_TV     DtapIE = 0x1C
	DESCRIPTION_OF_SECOND_CHANNEL_BEFORE_TIME_TV    DtapIE = 0x1D
	FREQUENCY_CHANNEL_SEQUENCE_BEFORE_TIME_TV       DtapIE = 0x1E
	MOBILE_ALLOCATION_BEFORE_TIME_TLV               DtapIE = 0x21
	VGCS_TARGET_MODE_INDICATION_TLV                 DtapIE = 0x01
	MULTI_RATE_CONFIGURATION_TLV                    DtapIE = 0x03
	VGCS_CIPHERING_PARAMETERS_TLV                   DtapIE = 0x04
	EXTENDED_TSC_SET_BEFORE_TIME_TV                 DtapIE = 0x6E
	FREQUENCY_SHORT_LIST_AFTER_TIME_TV              DtapIE = 0x02
	FREQUENCY_CHANNEL_SEQUENCE_AFTER_TIME_TV        DtapIE = 0x69
	REAL_TIME_DIFFERENCE_TLV                        DtapIE = 0x7B
	TIMING_ADVANCE_TV                               DtapIE = 0x7D
	FREQUENCY_SHORT_LIST_BEFORE_TIME_TV             DtapIE = 0x12
	DYNAMIC_ARFCN_MAPPING_TLV                       DtapIE = 0x76
	DEDICATED_SERVICE_INFORMATION_TV                DtapIE = 0x51
	MOBILE_OBSERVED_TIME_DIFF_TLV                   DtapIE = 0x77
	MOBILE_OBSERVED_TIME_DIFF_ON_HYPERFRAME_LVL_TLV DtapIE = 0x67
	BA_RANGE_TLV                                    DtapIE = 0x73
	GROUP_CHANNEL_DESCRIPTION_TLV                   DtapIE = 0x74
	BA_LIST_PREF_TLV                                DtapIE = 0x75
	GROUP_CHANNEL_DESCRIPTION_2_TLV                 DtapIE = 0x78
	TALKER_IDENTITY_TLV                             DtapIE = 0x79
	TALKER_PRIORITY_STATUS_TLV                      DtapIE = 0x7A
	BEARER_CAPABILITY_1_TLV                         DtapIE = 0x04 //CC embedded tags
	FACILITY_TLV                                    DtapIE = 0x1C
	PROGRESS_INDICATOR_TLV                          DtapIE = 0x1E
	SIGNAL_TV                                       DtapIE = 0x34
	CALLING_PARTY_BCD_NUMBER_TLV                    DtapIE = 0x5C
	CALLING_PARTY_SUB_ADDRESS_TLV                   DtapIE = 0x5D
	CALLED_PARTY_BCD_NUMBER_TLV                     DtapIE = 0x5E
	CALLED_PARTY_SUB_ADDRESS_TLV                    DtapIE = 0x6D
	REDIRECTING_PARTY_BCD_NUMBER_TLV                DtapIE = 0x74
	REDIRECTING_PARTY_SUB_ADDRESS_TLV               DtapIE = 0x75
	LOW_LAYER_COMPATIBILITY_I_TLV                   DtapIE = 0x7c
	HIGH_LAYER_COMPATIBILITY_I_TLV                  DtapIE = 0x7D
	USER_USER_TLV                                   DtapIE = 0x7E
	ALERT_TLV                                       DtapIE = 0x19
	NETWORK_CALL_CONTROL_CAPABILITIES_TLV           DtapIE = 0x2F
	CAUSE_OF_NO_CLI_TLV                             DtapIE = 0x3A
	BACKUP_BEARER_CAPABILITY_TLV                    DtapIE = 0x41
	ALLOWED_ACTIONS_CCBS_TLV                        DtapIE = 0x7B
	CAUSE_TLV                                       DtapIE = 0x08
	SS_VERSION_TLV                                  DtapIE = 0x7F
	CLIR_SUPPRESSION_T                              DtapIE = 0xA1
	CLIR_INVOCATION_T                               DtapIE = 0xA2
	CC_CAPABILITIES_TLV                             DtapIE = 0x15
	FACILITY_CCBS_ADVANCED_RECALL_ALIGNMENT_TLV     DtapIE = 0x1D
	FACILITY_RECALL_ALIGNMENT_NOT_ESSENTIAL_TLV     DtapIE = 0x1B
	STREAM_IDENTIFIER_TLV                           DtapIE = 0x2D
	SUPPORTED_CODECS_TLV                            DtapIE = 0x40
	REDIAL_T                                        DtapIE = 0xA3
	MODE_OF_CHANNEL_SET_7_TV                        DtapIE = 0x400 //duplicated iei in one PD
	VGCS_AMR_CONFIGURATION_TLV                      DtapIE = 0x401
	INDIVIDUAL_PRIORITIES_TLV                       DtapIE = 0x402
	UTRAN_FREQ_LIST_TLV                             DtapIE = 0x403
	CELL_SELECTION_INDICATOR_AFTER_RELEASE_TLV      DtapIE = 0x404
	MOBILE_IDENTITY_2_TLV                           DtapIE = 0x405
	MOBILE_IDENTITY_3_TLV                           DtapIE = 0x406
	LOW_LAYER_COMPATIBILITY_II_TLV                  DtapIE = 0x407
	BEARER_CAPABILITY_2_TLV                         DtapIE = 0x408
	HIGH_LAYER_COMPATIBILITY_II_TLV                 DtapIE = 0x409
	SECOND_CAUSE_TLV                                DtapIE = 0x41A
)

type IEFormat int

const (
	FormatUnsupported IEFormat = -1   // Unsupported element
	FormatT           IEFormat = iota // TAG
	FormatV                           // Value only (fixed length). Length is known in advance.
	FormatTV                          // Type + Value (fixed length, half-octet or full). Length is known.
	FormatLV                          // Length + Value (variable length). Length is read from the first byte.
	FormatTLV                         // Type + Length + Value (variable length). Length is read from the second byte.
)

/*
	Fixed length for V and TV.
	0 means half-octet for  V, there are always 2 consecutive nibbles
	-2 means The value is in the lower nibble, and the TV field length is 1 byte.
	Ignored for LV/TLV.
	SpecialHandling
	0 means nothing
	1 means that The high nibble determines which of the following conditional fields exists.
	2 means that This is a variable-length field in V-format.
*/

type IEDefinition struct {
	Format          IEFormat
	FixedLen        int
	Tag             DtapIE
	SpecialHandling int
}

func format(ie DtapIE, pd PD_Type) IEDefinition {
	switch pd {
	case PD_MM:
		return formatMM(ie)
	case PD_RR:
		return formatRR(ie)
	case PD_BCAST_CC:
		return formatBCCH(ie)
	case PD_CC:
		return formatCC(ie)
	case PD_SMS:
		return formatSMS(ie)
	case PD_GPRS_MMM:
		return formatGPRSMM(ie)
	case PD_GPRS_SMM:
		return formatGPRSSM(ie)
	case PD_LOC:
		return formatLOC(ie)
	case PD_GROUP_CC:
		return formatGroupCC(ie)
	case PD_EPS_SMM:
		return formatEPSSMM(ie)
	case PD_GTTP:
		return formatGTTP(ie)
	case PD_SS_NCL:
		return formatNCSS(ie)
	case PD_EXTEND:
		return formatExtend(ie)
	case PD_TEST:
		return formatTest(ie)
	case PD_EPS_MMM:
		return formatEPSMMM(ie)
	default:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	}
}

func formatMM(ie DtapIE) IEDefinition {
	switch ie {
	case PROTOCOL_DISC_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: PROTOCOL_DISC_V}
	case SKIP_IND_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: SKIP_IND_V}
	case MSG_TYPE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: MSG_TYPE_V}
	case CIPHERING_KEY_SEQUENCE_NUMBER_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CIPHERING_KEY_SEQUENCE_NUMBER_V}
	case SPARE_HALF_OCTET_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: SPARE_HALF_OCTET_V}
	case AUTH_PARAM_RAND_V:
		return IEDefinition{Format: FormatV, FixedLen: 16, Tag: AUTH_PARAM_RAND_V}
	case AUTH_RESP_PARAM_V:
		return IEDefinition{Format: FormatV, FixedLen: 4, Tag: AUTH_RESP_PARAM_V}
	case REJ_CAUSE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: REJ_CAUSE_V}
	case MOBILE_STATION_CLASSMARK_LV:
		return IEDefinition{Format: FormatLV, FixedLen: 4, Tag: MOBILE_STATION_CLASSMARK_LV}
	case MOBILE_IDENTITY_LV:
		return IEDefinition{Format: FormatLV, FixedLen: 0, Tag: MOBILE_IDENTITY_LV}
	case PD_AND_SAPI_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: PD_AND_SAPI_V}
	case CM_SERVICE_TYPE_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CM_SERVICE_TYPE_V}
	case IDENTITY_TYPE_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: IDENTITY_TYPE_V}
	case MOBILE_STATION_CLASSMARK_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: MOBILE_STATION_CLASSMARK_V}
	case LOCATION_UPDATING_TYPE_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: LOCATION_UPDATING_TYPE_V}
	case AUTH_PARAM_AUTN_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 18, Tag: AUTH_PARAM_AUTN_TLV}
	case AUTH_RESP_PARAM_EXT_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: AUTH_RESP_PARAM_EXT_TLV}
	case AUTH_FAILURE_PARAM_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 16, Tag: AUTH_FAILURE_PARAM_TLV}
	case LOCATION_AREA_ID_V:
		return IEDefinition{Format: FormatV, FixedLen: 5, Tag: LOCATION_AREA_ID_V}
	case LOCATION_AREA_ID_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 6, Tag: LOCATION_AREA_ID_TV}
	case DEVICE_PROPS_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case MM_TIMER_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: MM_TIMER_TLV}
	case PRIORITY_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x08}
	case ADDITIONAL_UPDATE_PARAMETERS_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0C}
	case P_TMSI_TYPE_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0E}
	case ROUTING_AREA_IDENTIFICATION_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 8, Tag: ROUTING_AREA_IDENTIFICATION_TLV}
	case P_TMSI_SIGNATURE_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 5, Tag: P_TMSI_SIGNATURE_TLV}
	case MOBILE_IDENTITY_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: MOBILE_IDENTITY_TLV}
	case FOLLOW_ON_PROCEED_T:
		return IEDefinition{Format: FormatT, FixedLen: 1, Tag: FOLLOW_ON_PROCEED_T}
	case CTS_PERMISSION_T:
		return IEDefinition{Format: FormatT, FixedLen: 1, Tag: CTS_PERMISSION_T}
	case EQUIVALENT_PLMNS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: EQUIVALENT_PLMNS_TLV}
	case EMERGENCY_NUMBER_LIST_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: EMERGENCY_NUMBER_LIST_TLV}
	case PER_MS_T3212_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: PER_MS_T3212_TLV}
	case NON_3GPP_NW_PROVIDED_POLICIES_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case MOBILE_STATION_CLASSMARK_FOR_UMTS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 5, Tag: MOBILE_STATION_CLASSMARK_FOR_UMTS_TLV}
	case MS_NETWORK_FEATURE_SUPPORT_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0E}
	case FULL_NAME_FOR_NETWORK_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: FULL_NAME_FOR_NETWORK_TLV}
	case SHORT_NAME_FOR_NETWORK_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: SHORT_NAME_FOR_NETWORK_TLV}
	case LOCAL_TIME_ZONE_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: LOCAL_TIME_ZONE_TV}
	case UNIVERSAL_TIME_AND_LOCAL_TIME_ZONE_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 8, Tag: UNIVERSAL_TIME_AND_LOCAL_TIME_ZONE_TV}
	case LSA_IDENTITY_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: LSA_IDENTITY_TLV}
	case NETWORK_DAYLIGHT_SAVING_TIME_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: NETWORK_DAYLIGHT_SAVING_TIME_TLV}
	default:
		return IEDefinition{Format: FormatUnsupported, FixedLen: -1}
	}
}

func formatRR(ie DtapIE) IEDefinition {
	switch ie {
	case PROTOCOL_DISC_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: PROTOCOL_DISC_V}
	case SKIP_IND_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: SKIP_IND_V}
	case MSG_TYPE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: MSG_TYPE_V}
	case PAGE_MODE_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: PAGE_MODE_V}
	case DEDICATED_MODE_OR_TBF_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: DEDICATED_MODE_OR_TBF_V, SpecialHandling: 1} // table 9.1.18.1 IMM ASSiGNMENT
	case CHANNEL_DESC_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: CHANNEL_DESC_V}
	case PACKET_CHANNEL_DESC_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: PACKET_CHANNEL_DESC_V}
	case REQUEST_REF_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: REQUEST_REF_V}
	case TIMING_ADVANCE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: TIMING_ADVANCE_V}
	case MOBILE_ALLOC_LV:
		return IEDefinition{Format: FormatLV, FixedLen: 0, Tag: MOBILE_ALLOC_LV}
	case STARTING_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 3, Tag: STARTING_TIME_TV}
	case IA_REST_OCTETS_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: IA_REST_OCTETS_V, SpecialHandling: 2}
	case EXTENDED_TSC_SET_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: EXTENDED_TSC_SET_TV}
	case FEATURE_INDICATOR_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: FEATURE_INDICATOR_V}
	case REQUEST_REF_1_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: REQUEST_REF_1_V}
	case WAIT_INDICATION_1_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: WAIT_INDICATION_1_V}
	case REQUEST_REF_2_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: REQUEST_REF_2_V}
	case WAIT_INDICATION_2_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: WAIT_INDICATION_2_V}
	case REQUEST_REF_3_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: REQUEST_REF_3_V}
	case WAIT_INDICATION_3_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: WAIT_INDICATION_3_V}
	case REQUEST_REF_4_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: REQUEST_REF_4_V}
	case WAIT_INDICATION_4_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: WAIT_INDICATION_4_V}
	case IAR_REST_OCTETS_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: IAR_REST_OCTETS_V}
	case CIPHERING_MODE_SETTING_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CIPHERING_MODE_SETTING_V}
	case CIPHER_RESPONSE_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CIPHER_RESPONSE_V}
	case ME_IDENTITY_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: ME_IDENTITY_TLV}
	case RR_CAUSE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: RR_CAUSE_V}
	case DESCRIPTION_OF_FIRST_CHANNEL_AFTER_TIME_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: DESCRIPTION_OF_FIRST_CHANNEL_AFTER_TIME_V}
	case POWER_COMMAND_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: POWER_COMMAND_V}
	case FREQUENCY_LIST_AFTER_TIME_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: FREQUENCY_LIST_AFTER_TIME_TLV}
	case CELL_CHANNEL_DESCRIPTION_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 17, Tag: CELL_CHANNEL_DESCRIPTION_TV}
	case DESCRIPTION_OF_MULTISLOT_CONFIGURATION_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: DESCRIPTION_OF_MULTISLOT_CONFIGURATION_TLV}
	case MODE_OF_FIRST_CHANNEL_CHANNEL_SET_1_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_FIRST_CHANNEL_CHANNEL_SET_1_TV}
	case MODE_OF_CHANNEL_SET_2_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_CHANNEL_SET_2_TV}
	case MODE_OF_CHANNEL_SET_3_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_CHANNEL_SET_3_TV}
	case MODE_OF_CHANNEL_SET_4_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_CHANNEL_SET_4_TV}
	case MODE_OF_CHANNEL_SET_5_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_CHANNEL_SET_5_TV}
	case MODE_OF_CHANNEL_SET_6_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_CHANNEL_SET_6_TV}
	case MODE_OF_CHANNEL_SET_7_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: 0x17}
	case MODE_OF_CHANNEL_SET_8_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_CHANNEL_SET_8_TV}
	case DESCRIPTION_OF_SECOND_CHANNEL_AFTER_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 4, Tag: DESCRIPTION_OF_SECOND_CHANNEL_AFTER_TIME_TV}
	case MODE_OF_SECOND_CHANNEL_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: MODE_OF_SECOND_CHANNEL_TV}
	case MOBILE_ALLOCATION_AFTER_TIME_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: MOBILE_ALLOCATION_AFTER_TIME_TLV}
	case FREQUENCY_LIST_BEFORE_TIME_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: FREQUENCY_LIST_BEFORE_TIME_TLV}
	case DESCRIPTION_OF_FIRST_CHANNEL_BEFORE_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 4, Tag: DESCRIPTION_OF_FIRST_CHANNEL_BEFORE_TIME_TV}
	case DESCRIPTION_OF_SECOND_CHANNEL_BEFORE_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 4, Tag: DESCRIPTION_OF_SECOND_CHANNEL_BEFORE_TIME_TV}
	case FREQUENCY_CHANNEL_SEQUENCE_BEFORE_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 10, Tag: FREQUENCY_CHANNEL_SEQUENCE_BEFORE_TIME_TV}
	case MOBILE_ALLOCATION_BEFORE_TIME_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: MOBILE_ALLOCATION_BEFORE_TIME_TLV}
	case CIPHER_MODE_SETTING_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x09}
	case VGCS_TARGET_MODE_INDICATION_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: VGCS_TARGET_MODE_INDICATION_TLV}
	case MULTI_RATE_CONFIGURATION_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: MULTI_RATE_CONFIGURATION_TLV}
	case VGCS_CIPHERING_PARAMETERS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: VGCS_CIPHERING_PARAMETERS_TLV}
	case EXTENDED_TSC_SET_BEFORE_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: EXTENDED_TSC_SET_BEFORE_TIME_TV}
	case CELL_DESCRIPTION_V:
		return IEDefinition{Format: FormatV, FixedLen: 2, Tag: CELL_DESCRIPTION_V}
	case HANDOVER_REFERENCE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: HANDOVER_REFERENCE_V}
	case POWER_COMMAND_AND_ACCESS_TYPE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: POWER_COMMAND_AND_ACCESS_TYPE_V}
	case SYNCHRONIZATION_INDICATION_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case FREQUENCY_SHORT_LIST_AFTER_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 10, Tag: FREQUENCY_SHORT_LIST_AFTER_TIME_TV}
	case FREQUENCY_CHANNEL_SEQUENCE_AFTER_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 10, Tag: FREQUENCY_CHANNEL_SEQUENCE_AFTER_TIME_TV}
	case REAL_TIME_DIFFERENCE_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: REAL_TIME_DIFFERENCE_TLV}
	case TIMING_ADVANCE_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: TIMING_ADVANCE_TV}
	case FREQUENCY_SHORT_LIST_BEFORE_TIME_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 10, Tag: FREQUENCY_SHORT_LIST_BEFORE_TIME_TV}
	case DYNAMIC_ARFCN_MAPPING_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: DYNAMIC_ARFCN_MAPPING_TLV}
	case DEDICATED_SERVICE_INFORMATION_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: DEDICATED_SERVICE_INFORMATION_TV}
	case PLMN_INDEX_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 1, Tag: 0x0A}
	case MOBILE_OBSERVED_TIME_DIFF_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: MOBILE_OBSERVED_TIME_DIFF_TLV}
	case MOBILE_OBSERVED_TIME_DIFF_ON_HYPERFRAME_LVL_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: MOBILE_OBSERVED_TIME_DIFF_ON_HYPERFRAME_LVL_TLV}
	case PS_CAUSE_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x09}
	case BA_RANGE_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: BA_RANGE_TLV}
	case GROUP_CHANNEL_DESCRIPTION_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: GROUP_CHANNEL_DESCRIPTION_TLV}
	case GROUP_CIPHER_KEY_NUMBER_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x08}
	case GPRS_RESUMPTION_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0C}
	case BA_LIST_PREF_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: BA_LIST_PREF_TLV}
	case UTRAN_FREQ_LIST_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x76}
	case CELL_SELECTION_INDICATOR_AFTER_RELEASE_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x77}
	case ENHANCED_DTM_CS_RELEASE_INDICATION_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0A}
	case GROUP_CHANNEL_DESCRIPTION_2_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 13, Tag: GROUP_CHANNEL_DESCRIPTION_2_TLV}
	case TALKER_IDENTITY_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: TALKER_IDENTITY_TLV}
	case TALKER_PRIORITY_STATUS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: TALKER_PRIORITY_STATUS_TLV}
	case VGCS_AMR_CONFIGURATION_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: 0x7B}
	case INDIVIDUAL_PRIORITIES_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x7C}
	case CHANNELS_NEEDED_FOR_MOBILES_1_AND_2_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CHANNELS_NEEDED_FOR_MOBILES_1_AND_2_V}
	case MOBILE_IDENTITY_1_LV:
		return IEDefinition{Format: FormatLV, FixedLen: 0, Tag: MOBILE_IDENTITY_1_LV}
	case MOBILE_IDENTITY_2_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x17}
	case P1_REST_OCTETS_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: P1_REST_OCTETS_V, SpecialHandling: 2}
	case MOBILE_IDENTITY_1_V:
		return IEDefinition{Format: FormatV, FixedLen: 4, Tag: MOBILE_IDENTITY_1_V}
	case MOBILE_IDENTITY_2_V:
		return IEDefinition{Format: FormatV, FixedLen: 4, Tag: MOBILE_IDENTITY_2_V}
	case MOBILE_IDENTITY_3_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x17}
	case P2_REST_OCTETS_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: P2_REST_OCTETS_V, SpecialHandling: 2}
	case MOBILE_IDENTITY_3_V:
		return IEDefinition{Format: FormatV, FixedLen: 4, Tag: MOBILE_IDENTITY_3_V}
	case MOBILE_IDENTITY_4_V:
		return IEDefinition{Format: FormatV, FixedLen: 4, Tag: MOBILE_IDENTITY_4_V}
	case P3_REST_OCTETS_V:
		return IEDefinition{Format: FormatV, FixedLen: 3, Tag: P3_REST_OCTETS_V}
	case CIPHERING_KEY_SEQUENCE_NUMBER_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CIPHERING_KEY_SEQUENCE_NUMBER_V}
	case SPARE_HALF_OCTET_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: SPARE_HALF_OCTET_V}
	case MOBILE_STATION_CLASSMARK_LV:
		return IEDefinition{Format: FormatLV, FixedLen: 4, Tag: MOBILE_STATION_CLASSMARK_LV}
	case MOBILE_IDENTITY_LV:
		return IEDefinition{Format: FormatLV, FixedLen: 0, Tag: MOBILE_IDENTITY_LV}
	case ADDITIONAL_UPDATE_PARAMETERS_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0C}
	default:
		return IEDefinition{Format: FormatUnsupported, FixedLen: -1}
	}
}

func formatCC(ie DtapIE) IEDefinition {
	switch ie {
	case PROTOCOL_DISC_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: PROTOCOL_DISC_V}
	case TRANSACTION_IDEN_V:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: TRANSACTION_IDEN_V}
	case MSG_TYPE_V:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: MSG_TYPE_V}
	case BC_REPEAT_INDICATOR_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case BEARER_CAPABILITY_1_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: BEARER_CAPABILITY_1_TLV}
	case BEARER_CAPABILITY_2_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x04}
	case FACILITY_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: FACILITY_TLV}
	case PROGRESS_INDICATOR_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: PROGRESS_INDICATOR_TLV}
	case SIGNAL_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x34}
	case CALLING_PARTY_BCD_NUMBER_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: CALLING_PARTY_BCD_NUMBER_TLV}
	case CALLING_PARTY_SUB_ADDRESS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: CALLING_PARTY_SUB_ADDRESS_TLV}
	case CALLED_PARTY_BCD_NUMBER_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: CALLED_PARTY_BCD_NUMBER_TLV}
	case CALLED_PARTY_SUB_ADDRESS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: CALLED_PARTY_SUB_ADDRESS_TLV}
	case REDIRECTING_PARTY_BCD_NUMBER_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: REDIRECTING_PARTY_BCD_NUMBER_TLV}
	case REDIRECTING_PARTY_SUB_ADDRESS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: REDIRECTING_PARTY_SUB_ADDRESS_TLV}
	case LLC_REPEAT_INDICATOR_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case LOW_LAYER_COMPATIBILITY_I_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: LOW_LAYER_COMPATIBILITY_I_TLV}
	case LOW_LAYER_COMPATIBILITY_II_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x7C}
	case HLC_REPEAT_INDICATOR_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case HIGH_LAYER_COMPATIBILITY_I_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: HIGH_LAYER_COMPATIBILITY_I_TLV}
	case HIGH_LAYER_COMPATIBILITY_II_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x7D}
	case USER_USER_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: USER_USER_TLV}
	case PRIORITY_TV:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x08}
	case ALERT_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: ALERT_TLV}
	case NETWORK_CALL_CONTROL_CAPABILITIES_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: NETWORK_CALL_CONTROL_CAPABILITIES_TLV}
	case CAUSE_OF_NO_CLI_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: CAUSE_OF_NO_CLI_TLV}
	case BACKUP_BEARER_CAPABILITY_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: BACKUP_BEARER_CAPABILITY_TLV}
	case CAUSE_LV:
		return IEDefinition{Format: FormatLV, FixedLen: 0, Tag: CAUSE_LV}
	case ALLOWED_ACTIONS_CCBS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: ALLOWED_ACTIONS_CCBS_TLV}
	case CAUSE_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: CAUSE_TLV}
	case SECOND_CAUSE_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: 0x08}
	case SS_VERSION_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: SS_VERSION_TLV}
	case CLIR_SUPPRESSION_T:
		return IEDefinition{Format: FormatT, FixedLen: 1, Tag: CLIR_SUPPRESSION_T}
	case CLIR_INVOCATION_T:
		return IEDefinition{Format: FormatT, FixedLen: 1, Tag: CLIR_INVOCATION_T}
	case CC_CAPABILITIES_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 4, Tag: CC_CAPABILITIES_TLV}
	case FACILITY_CCBS_ADVANCED_RECALL_ALIGNMENT_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: FACILITY_CCBS_ADVANCED_RECALL_ALIGNMENT_TLV}
	case FACILITY_RECALL_ALIGNMENT_NOT_ESSENTIAL_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: FACILITY_RECALL_ALIGNMENT_NOT_ESSENTIAL_TLV}
	case STREAM_IDENTIFIER_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: STREAM_IDENTIFIER_TLV}
	case SUPPORTED_CODECS_TLV:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: SUPPORTED_CODECS_TLV}
	case REDIAL_T:
		return IEDefinition{Format: FormatT, FixedLen: 1, Tag: REDIAL_T}
	default:
		return IEDefinition{Format: FormatUnsupported, FixedLen: -1}
	}
}

// ToDO
func formatBCCH(ie DtapIE) IEDefinition    { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatSMS(ie DtapIE) IEDefinition     { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatGPRSMM(ie DtapIE) IEDefinition  { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatGPRSSM(ie DtapIE) IEDefinition  { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatLOC(ie DtapIE) IEDefinition     { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatGroupCC(ie DtapIE) IEDefinition { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatEPSSMM(ie DtapIE) IEDefinition  { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatEPSMMM(ie DtapIE) IEDefinition  { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatGTTP(ie DtapIE) IEDefinition    { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatNCSS(ie DtapIE) IEDefinition    { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatExtend(ie DtapIE) IEDefinition  { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatTest(ie DtapIE) IEDefinition    { return IEDefinition{Format: FormatV, FixedLen: 0} }
