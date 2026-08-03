package dtap

func GetTagsOrder(pd PD_Type, msgType Msg_Type) []DtapIE {
    switch pd {
    case PD_MM:
        return getMMTagsOrder(msgType)
    case PD_RR:
        return getRRTagsOrder(msgType)
    case PD_BCAST_CC:
        return getBCCHTagsOrder(msgType)
    case PD_CC:
        return getCCTagsOrder(msgType)
    case PD_SMS:
        return getSMSTagsOrder(msgType)
    case PD_GPRS_MMM:
        return getGPRSMMTagsOrder(msgType)
    case PD_GPRS_SMM:
        return getGPRSSMTagsOrder(msgType)
    case PD_LOC:
        return getLOCTagsOrder(msgType)
    case PD_GROUP_CC:
        return getGroupCCTagsOrder(msgType)
    case PD_EPS_SMM:
        return getEPSSMMTagsOrder(msgType)
    case PD_GTTP:
        return getGTTTTagsOrder(msgType)
    case PD_SS_NCL:
        return getNCSSTagsOrder(msgType)
    case PD_EXTEND:
        return getExtendTagsOrder(msgType)
    case PD_TEST:
        return getTestTagsOrder(msgType)
    case PD_EPS_MMM:
        return getEPSMMMTagsOrder(msgType)
    default:
        return []DtapIE{}
    }
}

func getMMTagsOrder(msgType Msg_Type) []DtapIE {
    switch msgType {
    case MSG_MM_IMSI_DETACH_IND:
        return []DtapIE{
            MS_CLASSMARK_1,
            M_IDENTITY_1,
        }

    case MSG_MM_LOC_UPD_ACCEPT:
        return []DtapIE{
            LOC_AREA_ID,
            M_IDENTITY_2,
            FLLW_ON_PROC,
            CTS_PERM,
            PLMN_LST,
            EMER_NUM_LST,
            GPRS_TIM_3,
            NON_3GPP,
        }

    case MSG_MM_LOC_UPD_REJECT:
        return []DtapIE{
            REJ_CAUSE,
            MM_TIMER,
        }

    case MSG_MM_LOC_UPD_REQUEST:
        return []DtapIE{
            LOC_UPD_TYPE,
            CIPH_KEY_SEQ_NUM,
            LOC_AREA_ID,
            MS_CLASSMARK_1,
            M_IDENTITY_1,
            MS_CLASSMARK_UMTS,
            ADD_UPD_PARAMS,
            DEVICE_PROPS,
            MS_NET_FEAT_SUP,
        }

    case MSG_MM_AUTH_REQ:
        return []DtapIE{
            CIPH_KEY_SEQ_NUM,
            SPARE_HALF_OCT,
            AUTH_PARAM_RAND,
            AUTH_PARAM_AUTN,
        }

    case MSG_MM_AUTH_RESP:
        return []DtapIE{
            AUTH_RESP_PARAM,
            AUTH_RESP_PARAM_EXT,
        }

    case MSG_MM_AUTH_REJ:
        return []DtapIE{}

    case MSG_MM_AUTH_FAIL:
        return []DtapIE{
            REJ_CAUSE,
            AUTH_FAIL_PARAM,
        }

    case MSG_MM_ID_REQ:
        return []DtapIE{
            ID_TYPE,
            SPARE_HALF_OCT,
        }

    case MSG_MM_ID_RESP:
        return []DtapIE{
            M_IDENTITY_1,
            P_TMSI_TYPE,
            ROUT_AREA_ID_2,
            P_TMSI_SIGN_2,
        }

    case MSG_MM_TMSI_REALL_CMD:
        return []DtapIE{
            LOC_AREA_ID,
            M_IDENTITY_1,
        }

    case MSG_MM_TMSI_REALL_COMPL:
        return []DtapIE{}

    case MSG_MM_CM_SERV_ACC:
        return []DtapIE{}

    case MSG_MM_CM_SERV_REJ:
        return []DtapIE{
            REJ_CAUSE,
            MM_TIMER,
        }

    case MSG_MM_CM_SERV_ABORT:
        return []DtapIE{}

    case MSG_MM_CM_SERV_REQ:
        return []DtapIE{
            CM_SERVICE_TYPE,
            CIPH_KEY_SEQ_NUM,
            MS_CLASSMARK_2,
            M_IDENTITY_1,
            PRIOR_LVL,
            ADD_UPD_PARAMS,
            DEVICE_PROPS,
        }

    case MSG_MM_CM_SERV_PROMPT:
        return []DtapIE{PD_AND_SAPI}

    case MSG_MM_CM_REEST_REQ:
        return []DtapIE{
            CIPH_KEY_SEQ_NUM,
            SPARE_HALF_OCT,
            MS_CLASSMARK_2,
            M_IDENTITY_1,
            LOC_AREA_ID,
            DEVICE_PROPS,
        }

    case MSG_MM_ABORT:
        return []DtapIE{REJ_CAUSE}

    case MSG_MM_NULL:
        return []DtapIE{}
    
    case MSG_MM_STATUS:
        return []DtapIE{REJ_CAUSE}

    case MSG_MM_INFO:
        return  []DtapIE{
            FNAME_F_NET,
            SNAME_F_NET,
            TIME_ZONE,
            TIME_ZONE_AND_TIME,
            LSA_IDEN,
            DAY_SAVING_TIME,
        }

    default:
        return []DtapIE{}
    }
}

func getRRTagsOrder(msgType Msg_Type) []DtapIE {
    switch msgType {
    
    case MSG_RR_IMM_ASS:
        return []DtapIE{
            PAGE_MODE_V,
            DEDICATED_MODE_OR_TBF_V,
            CHANNEL_DESC_V,
            PACKET_CHANNEL_DESC_V,
            REQUEST_REF_V,
            TIMING_ADVANCE_V,
            MOBILE_ALLOC_LV,
            STARTING_TIME_TV,
            //IA_REST_OCTETS_V пока не поддерживается V переменной длины
            EXTENDED_TSC_SET_TV,
        }
    case MSG_RR_IMM_ASS_REJ:
        return []DtapIE{
            PAGE_MODE_V,
            FEATURE_INDICATOR_V,
            REQUEST_REF_1_V,
            WAIT_INDICATION_1_V,
            REQUEST_REF_2_V,
            WAIT_INDICATION_2_V,
            REQUEST_REF_3_V,
            WAIT_INDICATION_3_V,
            REQUEST_REF_4_V,
            WAIT_INDICATION_4_V,
            IAR_REST_OCTETS_V,
        }
    case MSG_RR_CIPH_M_CMD:
        return []DtapIE{
            CIPHERING_MODE_SETTING_V,
            CIPHER_RESPONSE_V,
        }
    case MSG_RR_CIPH_M_COMPL:
        return []DtapIE {
            ME_IDENTITY_TLV,
        }
    case MSG_RR_ASS_COMPL:
        return []DtapIE{
            RR_CAUSE_V,
        }
    case MSG_RR_ASS_FAIL:
        return []DtapIE {
            RR_CAUSE_V,
        }
    default:
        return []DtapIE{}
    }
}

//ToDO
func getBCCHTagsOrder(msgType Msg_Type) []DtapIE       { return nil }
func getCCTagsOrder(msgType Msg_Type) []DtapIE         { return nil }
func getSMSTagsOrder(msgType Msg_Type) []DtapIE        { return nil }
func getGPRSMMTagsOrder(msgType Msg_Type) []DtapIE     { return nil }
func getGPRSSMTagsOrder(msgType Msg_Type) []DtapIE     { return nil }
func getLOCTagsOrder(msgType Msg_Type) []DtapIE        { return nil }
func getGroupCCTagsOrder(msgType Msg_Type) []DtapIE    { return nil }
func getEPSSMMTagsOrder(msgType Msg_Type) []DtapIE     { return nil }
func getEPSMMMTagsOrder(msgType Msg_Type) []DtapIE     { return nil }
func getGTTTTagsOrder(msgType Msg_Type) []DtapIE       { return nil }
func getNCSSTagsOrder(msgType Msg_Type) []DtapIE       { return nil }
func getExtendTagsOrder(msgType Msg_Type) []DtapIE     { return nil }
func getTestTagsOrder(msgType Msg_Type) []DtapIE       { return nil } 