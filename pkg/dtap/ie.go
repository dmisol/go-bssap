package dtap

type DtapIE uint16

//Главное чтобы в рамках 1 PD не повторялись IE теги
//для одного и того же type могут быть разный FORMAT и разные IEI поэтому для каждого делаю тег по имени в RR. 
//В MM в основном по типу.

const (
	//general created Tags
	PROTOCOL_DISC							DtapIE = 0x100
	SKIP_IND								DtapIE = 0x101
	MSG_TYPE								DtapIE = 0x102
	CIPH_KEY_SEQ_NUM						DtapIE = 0x103
	SPARE_HALF_OCT							DtapIE = 0x104
	AUTH_PARAM_RAND							DtapIE = 0x105
	AUTH_RESP_PARAM							DtapIE = 0x106
	REJ_CAUSE								DtapIE = 0x107
	MS_CLASSMARK_2							DtapIE = 0x108
	M_IDENTITY_1							DtapIE = 0x109
	PD_AND_SAPI								DtapIE = 0x10A
	CM_SERVICE_TYPE							DtapIE = 0x10B
	ID_TYPE									DtapIE = 0x10C
	MS_CLASSMARK_1							DtapIE = 0x10D
	LOC_UPD_TYPE							DtapIE = 0x10E
	LOC_AREA_ID								DtapIE = 0x10F //table 9.2.15
	MS_NET_FEAT_SUP							DtapIE = 0x110
	NON_3GPP								DtapIE = 0x111
	P_TMSI_TYPE								DtapIE = 0x112
	DEVICE_PROPS							DtapIE = 0x113
	PAGE_MODE_V                             DtapIE = 0x114
	DEDICATED_MODE_OR_TBF_V                 DtapIE = 0x115
	CHANNEL_DESC_V                          DtapIE = 0x116
	PACKET_CHANNEL_DESC_V                   DtapIE = 0x117
	REQUEST_REF_V                           DtapIE = 0x118
	TIMING_ADVANCE_V                        DtapIE = 0x119
	MOBILE_ALLOC_LV                         DtapIE = 0x11A
	IA_REST_OCTETS_V                        DtapIE = 0x11B
	FEATURE_INDICATOR_V                     DtapIE = 0x11C
	REQUEST_REF_1_V                         DtapIE = 0x11D
	WAIT_INDICATION_1_V                     DtapIE = 0x11E
	REQUEST_REF_2_V                         DtapIE = 0x11F
	WAIT_INDICATION_2_V                     DtapIE = 0x120
	REQUEST_REF_3_V                         DtapIE = 0x121
	WAIT_INDICATION_3_V                     DtapIE = 0x122
	REQUEST_REF_4_V                         DtapIE = 0x123
	WAIT_INDICATION_4_V                     DtapIE = 0x124
	IAR_REST_OCTETS_V                       DtapIE = 0x125
	CIPHERING_MODE_SETTING_V                DtapIE = 0x126
	CIPHER_RESPONSE_V                       DtapIE = 0x127
	RR_CAUSE_V                              DtapIE = 0x128

	//MM
	//embedded tags
	AUTH_PARAM_AUTN							DtapIE = 0x20
	AUTH_RESP_PARAM_EXT						DtapIE = 0x21
	AUTH_FAIL_PARAM							DtapIE = 0x22
	LOC_AREA_ID_2							DtapIE = 0x13 //table 9.2.5
	MM_TIMER								DtapIE = 0x36
	PRIOR_LVL								DtapIE = 0x08
	ADD_UPD_PARAMS							DtapIE = 0x0C
	ROUT_AREA_ID_2							DtapIE = 0x1B		
	P_TMSI_SIGN_2							DtapIE = 0x19
	M_IDENTITY_2							DtapIE = 0x17
	FLLW_ON_PROC							DtapIE = 0xA1
	CTS_PERM								DtapIE = 0xA2
	PLMN_LST								DtapIE = 0x4A
	EMER_NUM_LST							DtapIE = 0x34
	GPRS_TIM_3								DtapIE = 0x35
	MS_CLASSMARK_UMTS						DtapIE = 0x33
	FNAME_F_NET								DtapIE = 0x43
	SNAME_F_NET								DtapIE = 0x45
	TIME_ZONE								DtapIE = 0x46
	TIME_ZONE_AND_TIME						DtapIE = 0x47
	LSA_IDEN								DtapIE = 0x48
	DAY_SAVING_TIME							DtapIE = 0x49
	//RR
	//embedded tags
	STARTING_TIME_TV                        DtapIE = 0x7C
	EXTENDED_TSC_SET_TV                     DtapIE = 0x6D
	ME_IDENTITY_TLV                         DtapIE = 0x17	
)

type IEFormat int

const (
    FormatUnsupported 	IEFormat = -1	// Unsupported element
	FormatT				IEFormat = iota	// TAG
    FormatV           				 	// Value only (fixed length). Length is known in advance.
    FormatTV                          	// Type + Value (fixed length, half-octet or full). Length is known.
    FormatLV                          	// Length + Value (variable length). Length is read from the first byte.
    FormatTLV                         	// Type + Length + Value (variable length). Length is read from the second byte.
)

/*	
	Fixed length for V and TV. 
	0 means half-octet for  V, there are always 2 consecutive nibbles 
	Ignored for LV/TLV.	
	SpecialHandling
	0 means nothing
	1 means that The high nibble determines which of the following conditional fields exists.
*/

type IEDefinition struct {
	Format   IEFormat
	FixedLen int			
	Tag DtapIE
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
	case PROTOCOL_DISC:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: PROTOCOL_DISC}
	case SKIP_IND:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: SKIP_IND}
	case MSG_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: MSG_TYPE}
	case CIPH_KEY_SEQ_NUM:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CIPH_KEY_SEQ_NUM}
	case SPARE_HALF_OCT:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: SPARE_HALF_OCT}
	case AUTH_PARAM_RAND:
		return IEDefinition{Format: FormatV, FixedLen: 16, Tag: AUTH_PARAM_RAND}
	case AUTH_RESP_PARAM:
		return IEDefinition{Format: FormatV, FixedLen: 4, Tag: AUTH_RESP_PARAM}
	case REJ_CAUSE:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: REJ_CAUSE}
	case MS_CLASSMARK_2:
		return IEDefinition{Format: FormatLV, FixedLen: 4, Tag: MS_CLASSMARK_2}
	case M_IDENTITY_1:
		return IEDefinition{Format: FormatLV, FixedLen: 0, Tag: M_IDENTITY_1}
	case PD_AND_SAPI:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: PD_AND_SAPI}
	case CM_SERVICE_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: CM_SERVICE_TYPE}
	case ID_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: ID_TYPE}
	case MS_CLASSMARK_1:
		return IEDefinition{Format: FormatV, FixedLen: 1, Tag: MS_CLASSMARK_1}
	case LOC_UPD_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 0, Tag: LOC_UPD_TYPE}
	case AUTH_PARAM_AUTN:
		return IEDefinition{Format: FormatTLV, FixedLen: 18, Tag: AUTH_PARAM_AUTN}
	case AUTH_RESP_PARAM_EXT:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: AUTH_RESP_PARAM_EXT}
	case AUTH_FAIL_PARAM:
		return IEDefinition{Format: FormatTLV, FixedLen: 16, Tag: AUTH_FAIL_PARAM}
	case LOC_AREA_ID:
		return IEDefinition{Format: FormatV, FixedLen: 5, Tag: LOC_AREA_ID}
	case LOC_AREA_ID_2:
		return IEDefinition{Format: FormatTV, FixedLen: 6, Tag: LOC_AREA_ID_2}
	case DEVICE_PROPS:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case MM_TIMER:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: MM_TIMER}
	case PRIOR_LVL:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x08}
	case ADD_UPD_PARAMS:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0C}
	case P_TMSI_TYPE:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0E}
	case ROUT_AREA_ID_2:
		return IEDefinition{Format: FormatTLV, FixedLen: 8, Tag: ROUT_AREA_ID_2}
	case P_TMSI_SIGN_2:
		return IEDefinition{Format: FormatTLV, FixedLen: 5, Tag: P_TMSI_SIGN_2}
	case M_IDENTITY_2:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: M_IDENTITY_2}
	case FLLW_ON_PROC:
		return IEDefinition{Format: FormatT, FixedLen: 1, Tag: FLLW_ON_PROC}
	case CTS_PERM:
		return IEDefinition{Format: FormatT, FixedLen: 1, Tag: CTS_PERM}
	case PLMN_LST:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: PLMN_LST}
	case EMER_NUM_LST:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: EMER_NUM_LST}
	case GPRS_TIM_3:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: GPRS_TIM_3}
	case NON_3GPP:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0D}
	case MS_CLASSMARK_UMTS:
		return IEDefinition{Format: FormatTLV, FixedLen: 5, Tag: MS_CLASSMARK_UMTS}
	case MS_NET_FEAT_SUP:
		return IEDefinition{Format: FormatTV, FixedLen: -2, Tag: 0x0E}
	case FNAME_F_NET:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: FNAME_F_NET}
	case SNAME_F_NET:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: SNAME_F_NET}
	case TIME_ZONE:
		return IEDefinition{Format: FormatTV, FixedLen: 2, Tag: TIME_ZONE}
	case TIME_ZONE_AND_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 8, Tag: TIME_ZONE_AND_TIME}
	case LSA_IDEN:
		return IEDefinition{Format: FormatTLV, FixedLen: 0, Tag: LSA_IDEN}
	case DAY_SAVING_TIME:
		return IEDefinition{Format: FormatTLV, FixedLen: 3, Tag: DAY_SAVING_TIME}
	default:
		return IEDefinition{Format: FormatUnsupported, FixedLen: -1}
	}
}

func formatRR(ie DtapIE) IEDefinition {
    switch ie {
	case PROTOCOL_DISC:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case SKIP_IND:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case MSG_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 1}
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
        return IEDefinition{Format: FormatV, FixedLen: 0, Tag: IA_REST_OCTETS_V}
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
	default:
		return IEDefinition{Format: FormatUnsupported, FixedLen: -1}
    }
}


//ToDO
func formatBCCH(ie DtapIE) IEDefinition       { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatCC(ie DtapIE) IEDefinition         { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatSMS(ie DtapIE) IEDefinition        { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatGPRSMM(ie DtapIE) IEDefinition     { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatGPRSSM(ie DtapIE) IEDefinition     { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatLOC(ie DtapIE) IEDefinition        { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatGroupCC(ie DtapIE) IEDefinition    { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatEPSSMM(ie DtapIE) IEDefinition     { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatEPSMMM(ie DtapIE) IEDefinition     { return IEDefinition{Format: FormatV, FixedLen: 0} } 
func formatGTTP(ie DtapIE) IEDefinition       { return IEDefinition{Format: FormatV, FixedLen: 0} } 
func formatNCSS(ie DtapIE) IEDefinition       { return IEDefinition{Format: FormatV, FixedLen: 0} }
func formatExtend(ie DtapIE) IEDefinition     { return IEDefinition{Format: FormatV, FixedLen: 0} } 
func formatTest(ie DtapIE) IEDefinition       { return IEDefinition{Format: FormatV, FixedLen: 0} } 