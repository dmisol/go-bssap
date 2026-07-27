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

	//embedded tags
	AUTH_PARAM_AUTN					DtapIE = 0x20
	AUTH_RESP_PARAM_EXT				DtapIE = 0x21
	AUTH_FAIL_PARAM					DtapIE = 0x22
	LOC_AREA_ID						DtapIE = 0x13
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
