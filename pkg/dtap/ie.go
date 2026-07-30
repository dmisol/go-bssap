package dtap

type DtapIE uint16

//Главное чтобы в рамках 1 PD не повторялись IE теги
//для одного и того же type могут быть разный FORMAT и разные IEI поэтому для каждого делаю тег по имени в RR. В MM в основном по типу.

const (
	//created tags
	//MM
	PROTOCOL_DISC					DtapIE = 0x181
	SKIP_IND						DtapIE = 0x182
	MSG_TYPE						DtapIE = 0x183
	CIPH_KEY_SEQ_NUM				DtapIE = 0x184
	SPARE_HALF_OCT					DtapIE = 0x185
	AUTH_PARAM_RAND					DtapIE = 0x186
	AUTH_RESP_PARAM					DtapIE = 0x187
	REJ_CAUSE						DtapIE = 0x188
	MS_CLASSMARK_2					DtapIE = 0x189
	M_IDENTITY_1					DtapIE = 0x18A
	PD_AND_SAPI						DtapIE = 0x18B
	CM_SERVICE_TYPE					DtapIE = 0x18C
	ID_TYPE							DtapIE = 0x18D
	MS_CLASSMARK_1					DtapIE = 0x18E
	LOC_UPD_TYPE					DtapIE = 0x18F
	LOC_AREA_ID						DtapIE = 0x190 //table 9.2.15

	//RR
	CHANNEL_DESC					DtapIE = 0x181

	//ASSIGNMENT CMD
	DESC_OF_THE_F_CH_AFTER_TIME		DtapIE = 0x182
	POWER_CMD						DtapIE = 0x183

	//ASSIGNMENT COMPLETE
	RR_CAUSE						DtapIE = 0x184

	//CHANNEL MODE MODIFY
	CHANNEL_MODE					DtapIE = 0x185

	//IMMEDIATE ASSIGNMENT
	PAGE_MODE              			DtapIE = 0x186
	DEDICATED_MODE_OR_TBF  			DtapIE = 0x187
	PACKET_CH_DESC         			DtapIE = 0x188
	REQ_REF                			DtapIE = 0x189
	TIMING_ADVANCE         			DtapIE = 0x18A
	MOBILE_ALLOC_2           		DtapIE = 0x18B
	IA_REST_OCTETS         			DtapIE = 0x18C

	//IMMEDIATE ASSIGNMENT REJECT
	FEATURE_IND       				DtapIE = 0x18D
	REQ_REF_1         				DtapIE = 0x18E
	WAIT_IND_1        				DtapIE = 0x18F
	REQ_REF_2         				DtapIE = 0x190
	WAIT_IND_2        				DtapIE = 0x191
	REQ_REF_3         				DtapIE = 0x192
	WAIT_IND_3        				DtapIE = 0x193
	REQ_REF_4         				DtapIE = 0x194
	WAIT_IND_4        				DtapIE = 0x195
	IAR_REST_OCTETS   				DtapIE = 0x196

	//CIPHERING MODE COMMAND
	CIPH_MODE_SET_V					DtapIE = 0x197
	CIPH_RESP						DtapIE = 0x198

	//HANDOVER COMMAND
	CELL_DESC                 		DtapIE = 0x199
	DESC_OF_FIRST_CH_AFTER    		DtapIE = 0x19A
	HANDOVER_REF              		DtapIE = 0x19B
	POWER_CMD_AND_ACCESS_TYPE 		DtapIE = 0x19C

	//embedded tags
	//MM
	AUTH_PARAM_AUTN					DtapIE = 0x20
	AUTH_RESP_PARAM_EXT				DtapIE = 0x21
	AUTH_FAIL_PARAM					DtapIE = 0x22
	LOC_AREA_ID_2					DtapIE = 0x13 //table 9.2.5
	DEVICE_PROPS					DtapIE = 0x0D		//тут аккуратно значение в младшем полубайте а IEI как D кодируется в старшем и у еще одного как D.
	MM_TIMER						DtapIE = 0x36
	PRIOR_LVL						DtapIE = 0x08
	ADD_UPD_PARAMS					DtapIE = 0x0C
	P_TMSI_TYPE						DtapIE = 0x0E		//тут аккуратно
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

	//RR
	MOBILE_ALLOC					DtapIE = 0x72
	START_TIME						DtapIE = 0x7C
	EXTEND_TSC_S					DtapIE = 0x6D
	//ASSIGNMENT CMD
	FREQ_LST_AFTER_TIME				DtapIE = 0x05
	CELL_CH_DESC					DtapIE = 0x62
	DESC_OF_THE_MULT_CONF			DtapIE = 0x10
	MODE_OF_CH_SET_1				DtapIE = 0x63
	MODE_OF_CH_SET_2				DtapIE = 0x11
	MODE_OF_CH_SET_3				DtapIE = 0x13
	MODE_OF_CH_SET_4				DtapIE = 0x14
	MODE_OF_CH_SET_5				DtapIE = 0x15
	MODE_OF_CH_SET_6				DtapIE = 0x16
	MODE_OF_CH_SET_7				DtapIE = 0x17
	MODE_OF_CH_SET_8				DtapIE = 0x18
	DESC_OF_THE_SCH					DtapIE = 0x64
	MODE_OF_THE_SCH					DtapIE = 0x66
	FREQ_LST_BEF_TIME				DtapIE = 0x19
	DESC_O_T_FIRST_CH_BEF_TIME		DtapIE = 0x1C
	DESC_O_T_SEC_CH_BEF_TIME		DtapIE = 0x1D
	FREQ_CH_SEQ_BEF_TIME			DtapIE = 0x1E
	MOB_ALLOC_BEFORE_TIME			DtapIE = 0x21
	CIPH_MODE_SET					DtapIE = 0x09
	VGCS_TARGET_MODE_IND 			DtapIE = 0x01
	MULTI_RATE_CONF					DtapIE = 0x03
	VGCS_CIPH_PARAMS				DtapIE = 0x04
	EXTEND_TSC_S_BEFORE_TIME		DtapIE = 0x6E
	//HANDOVER COMMAND
	SYNC_IND                  		DtapIE = 0x0D
	FREQ_CH_SEQ_AFTER_TIME    		DtapIE = 0x69
	FREQ_SHORT_LST_AFTER_TIME       DtapIE = 0x02
	REAL_TIME_DIFF           		DtapIE = 0x7B
	TIMING_ADVANCE_TV           	DtapIE = 0x7D
	FREQ_SHORT_LST_BEF_TIME 		DtapIE = 0x12
	DESC_O_T_SEC_CH_BEF      		DtapIE = 0x1D
	DYNAMIC_ARFCN_MAPPING    		DtapIE = 0x76
	DEDICATED_SERV_INFO      		DtapIE = 0x51
	PLMN_INDEX              		DtapIE = 0x0A
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

type IEDefinition struct {
	Format   IEFormat
	FixedLen int // Fixed length for V and TV. 0 means half-octet for TV OR V. Ignored for LV/TLV.
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

func formatRR(ie DtapIE) IEDefinition {
    switch ie {
	case CHANNEL_DESC:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case MOBILE_ALLOC:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case START_TIME:
		return  IEDefinition{Format: FormatTV, FixedLen: 3}
	case EXTEND_TSC_S:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case DESC_OF_THE_F_CH_AFTER_TIME:
		return  IEDefinition{Format: FormatV, FixedLen: 3}
	case POWER_CMD:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case FREQ_LST_AFTER_TIME:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case CELL_CH_DESC:
		return IEDefinition{Format: FormatTV, FixedLen: 17}
	case DESC_OF_THE_MULT_CONF:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case MODE_OF_CH_SET_1:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case MODE_OF_CH_SET_2:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case MODE_OF_CH_SET_3:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case MODE_OF_CH_SET_4:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case MODE_OF_CH_SET_5:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case MODE_OF_CH_SET_6:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case MODE_OF_CH_SET_7:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case MODE_OF_CH_SET_8:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case DESC_OF_THE_SCH:
		return IEDefinition{Format: FormatTV, FixedLen: 4}
	case MODE_OF_THE_SCH:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case FREQ_LST_BEF_TIME:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case DESC_O_T_FIRST_CH_BEF_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 4}
	case DESC_O_T_SEC_CH_BEF_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 4}
	case FREQ_CH_SEQ_BEF_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 10}
	case MOB_ALLOC_BEFORE_TIME:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case CIPH_MODE_SET:
		return IEDefinition{Format: FormatTV, FixedLen: 1}
	case VGCS_TARGET_MODE_IND:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case MULTI_RATE_CONF:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case VGCS_CIPH_PARAMS:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case EXTEND_TSC_S_BEFORE_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case RR_CAUSE:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case CHANNEL_MODE:
		return  IEDefinition{Format: FormatV, FixedLen: 1}

	case PAGE_MODE:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case DEDICATED_MODE_OR_TBF:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case PACKET_CH_DESC:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case REQ_REF:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case TIMING_ADVANCE:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case MOBILE_ALLOC_2:
		return IEDefinition{Format: FormatLV, FixedLen: 0}
	case IA_REST_OCTETS:
		return IEDefinition{Format: FormatV, FixedLen: 0} //!! на данный момент не работает тк нет длины и нужно отдельно парсить стр 400
	case FEATURE_IND:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case REQ_REF_1:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case WAIT_IND_1:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case REQ_REF_2:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case WAIT_IND_2:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case REQ_REF_3:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case WAIT_IND_3:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case REQ_REF_4:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case WAIT_IND_4:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case IAR_REST_OCTETS:
		return IEDefinition{Format: FormatV, FixedLen: 3}
	case CIPH_MODE_SET_V:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case CIPH_RESP:
		return IEDefinition{Format: FormatV, FixedLen: 0}
	case CELL_DESC:
		return  IEDefinition{Format: FormatV, FixedLen: 2}
	case DESC_OF_FIRST_CH_AFTER:
		return  IEDefinition{Format: FormatV, FixedLen: 3}
	case HANDOVER_REF:
		return IEDefinition{Format: FormatV, FixedLen: 1}
	case POWER_CMD_AND_ACCESS_TYPE:
		return  IEDefinition{Format: FormatV, FixedLen: 1}
	case SYNC_IND:
		return IEDefinition{Format: FormatTV, FixedLen: 1}
	case FREQ_CH_SEQ_AFTER_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 10}
	case FREQ_SHORT_LST_AFTER_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 10}
	case REAL_TIME_DIFF:
		return IEDefinition{Format: FormatTLV, FixedLen: 3}
	case TIMING_ADVANCE_TV:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case FREQ_SHORT_LST_BEF_TIME:
		return IEDefinition{Format: FormatTV, FixedLen: 10}
	case DYNAMIC_ARFCN_MAPPING:
		return IEDefinition{Format: FormatTLV, FixedLen: 0}
	case DEDICATED_SERV_INFO:
		return IEDefinition{Format: FormatTV, FixedLen: 2}
	case PLMN_INDEX:
		return IEDefinition{Format: FormatTV, FixedLen: 1}
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