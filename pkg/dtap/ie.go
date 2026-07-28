package dtap

type DtapIE byte

const (

	//created tags
	PROTOCOL_DISC					DtapIE = 0x81
	SKIP_IND						DtapIE	= 0x82
	MSG_TYPE						DtapIE = 0x83
	CIPH_KEY_SEQ_NUM				DtapIE = 0x84
	SPARE_HALF_OCT					DtapIE = 0x85
	AUTH_PARAM_RAND					DtapIE = 0x86
	AUTH_RESP_PARAM					DtapIE = 0x87
	REJ_CAUSE						DtapIE = 0x88
	MS_CLASSMARK_2					DtapIE = 0x89
	M_IDENTITY_1					DtapIE = 0x8A
	PD_AND_SAPI						DtapIE = 0x8B
	CM_SERVICE_TYPE					DtapIE = 0x8C
	ID_TYPE							DtapIE = 0x8D
	MS_CLASSMARK_1					DtapIE = 0x8E
	LOC_UPD_TYPE					DtapIE = 0x8F
	LOC_AREA_ID						DtapIE = 0x90 //table 9.2.15

	//embedded tags
	AUTH_PARAM_AUTN					DtapIE = 0x20
	AUTH_RESP_PARAM_EXT				DtapIE = 0x21
	AUTH_FAIL_PARAM					DtapIE = 0x22
	LOC_AREA_ID_2					DtapIE = 0x13 //table 9.2.5
	DEVICE_PROPS					DtapIE = 0x0D
	MM_TIMER						DtapIE = 0x36
	PRIOR_LVL						DtapIE = 0x08
	ADD_UPD_PARAMS					DtapIE = 0x0C
	P_TMSI_TYPE						DtapIE = 0x0E
	ROUT_AREA_ID_2					DtapIE = 0x1B		
	P_TMSI_SIGN_2					DtapIE = 0x19
	M_IDENTITY_2					DtapIE = 0x17
	FLLW_ON_PROC					DtapIE = 0xA1
	CTS_PERM						DtapIE = 0xA2
	PLMN_LST						DtapIE = 0x4A
	EMER_NUM_LST					DtapIE = 0x34
	GPRS_TIM_3						DtapIE = 0x35
	NON_3GPP						DtapIE = 0xD1		//тут аккуратно
	MS_CLASSMARK_UMTS				DtapIE = 0x33
	MS_NET_FEAT_SUP					DtapIE = 0xE1		//тут аккуратно	
	FNAME_F_NET						DtapIE = 0x43
	SNAME_F_NET						DtapIE = 0x45
	TIME_ZONE						DtapIE = 0x46
	TIME_ZONE_AND_TIME				DtapIE = 0x47
	LSA_IDEN						DtapIE = 0x48
	DAY_SAVING_TIME					DtapIE = 0x49
)

// format returns a length of the Information Element (IE)
// 1. If length > 0, it means that IE has FIXED length and V format
// 2. If length = 0, it means that IE has VARIABLE length and LV format
// 3. If length = -1, it means that IE is unsupported
// 4. If length = -2, it means that IE is half byte and has TV format
// 5. If length = -3, it means that IE has VARIABLE (or FIXED) size and has TLV format

type IEFormat int

const (
    FormatUnsupported 	IEFormat = -1	// Unsupported element
	FormatT				IEFormat = iota	// TAG
    FormatV           				 	// Value only (fixed length). Length is known in advance.
    FormatTV                          	// Type + Value (fixed length, half-octet or full). Length is known.
    FormatLV                          	// Length + Value (variable length). Length is read from the first byte.
    FormatTLV                         	// Type + Length + Value (variable length). Length is read from the second byte.
)

type IEDefinition struct {
	Format   IEFormat
	FixedLen int // Fixed length for V and TV. 0 means half-octet for TV OR V. Ignored for LV/TLV.
}

func (ie DtapIE) format() IEDefinition {
	switch ie {
	case PROTOCOL_DISC:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case SKIP_IND:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case MSG_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case CIPH_KEY_SEQ_NUM:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case SPARE_HALF_OCT:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case AUTH_PARAM_RAND:
		return IEDefinition{Format: FormatV, FixedLen: 16}
	case AUTH_RESP_PARAM:
		return IEDefinition{Format: FormatV, FixedLen: 4}
	case REJ_CAUSE:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case MS_CLASSMARK_2:
		return IEDefinition{Format: FormatLV, FixedLen: 4}
	case M_IDENTITY_1:
		return IEDefinition{Format: FormatLV, FixedLen: 0}
	case PD_AND_SAPI:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case CM_SERVICE_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case ID_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case MS_CLASSMARK_1:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case LOC_UPD_TYPE:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case AUTH_PARAM_AUTN:
		return IEDefinition{Format: FormatTLV, FixedLen: 18}
	case AUTH_RESP_PARAM_EXT:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case AUTH_FAIL_PARAM:
		return IEDefinition{Format: FormatTLV, FixedLen: 16}
	case LOC_AREA_ID:
		return IEDefinition{Format: FormatV, FixedLen: 5}
	case LOC_AREA_ID_2:
		return IEDefinition{Format: FormatTV, FixedLen: 6}
	case DEVICE_PROPS:
		return IEDefinition{Format: FormatTV, FixedLen: 0}
	case MM_TIMER:
		return IEDefinition{Format: FormatTLV, FixedLen: 3}
	case PRIOR_LVL:
		return IEDefinition{Format: FormatTV, FixedLen: 0}
	case ADD_UPD_PARAMS:
		return IEDefinition{Format: FormatTV, FixedLen: 0}
	case P_TMSI_TYPE:
		return IEDefinition{Format: FormatTV, FixedLen: 0}
	case ROUT_AREA_ID_2:
		return IEDefinition{Format: FormatTLV, FixedLen: 8}
	case P_TMSI_SIGN_2:
		return IEDefinition{Format: FormatTLV, FixedLen: 5}
	case M_IDENTITY_2:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case FLLW_ON_PROC:
		return IEDefinition{Format: FormatT, FixedLen: 1}
	case CTS_PERM:
		return IEDefinition{Format: FormatT, FixedLen: 1}
	case PLMN_LST:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case EMER_NUM_LST:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case GPRS_TIM_3:
		return IEDefinition{Format: FormatTLV, FixedLen: 3}
	case NON_3GPP:
		return IEDefinition{Format: FormatTV, FixedLen: 0}
	case MS_CLASSMARK_UMTS:
		return IEDefinition{Format: FormatTLV, FixedLen: 5}
	case MS_NET_FEAT_SUP:
		return IEDefinition{Format: FormatTV, FixedLen: 0}
	case FNAME_F_NET:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case SNAME_F_NET:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case TIME_ZONE:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case TIME_ZONE_AND_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 8}
	case LSA_IDEN:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case DAY_SAVING_TIME:
		return IEDefinition{Format: FormatTLV, FixedLen: 3}
	default:
		return IEDefinition{Format: FormatUnsupported, FixedLen: -1}
	}
}