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
    // Channel Management
    case MSG_RR_INIT_REQ:
        return []DtapIE{}
    case MSG_RR_ADD_ASS:
        return []DtapIE{
            CHANNEL_DESC,
            MOBILE_ALLOC,
            START_TIME,
            EXTEND_TSC_S,
        }
    case MSG_RR_IMM_ASS:
        return []DtapIE{}
    case MSG_RR_IMM_ASS_EXT:
        return []DtapIE{}
    case MSG_RR_IMM_ASS_REJ:
        return []DtapIE{}
    case MSG_RR_DTM_ASS_FAIL:
        return []DtapIE{}
    case MSG_RR_DTM_REJECT:
        return []DtapIE{}
    case MSG_RR_DTM_REQUEST:
        return []DtapIE{}
    case MSG_RR_PACKET_ASS:
        return []DtapIE{}

    // Ciphering
    case MSG_RR_CIPH_M_CMD:
        return []DtapIE{}
    case MSG_RR_CIPH_M_COMPL:
        return []DtapIE{}

    // Configuration Change
    case MSG_RR_CFG_CHG_CMD:
        return []DtapIE{}
    case MSG_RR_CFG_CHG_ACK:
        return []DtapIE{}
    case MSG_RR_CFG_CHG_REJ:
        return []DtapIE{}

    // Assignment & Handover
    case MSG_RR_ASS_CMD:
        return []DtapIE{}
    case MSG_RR_ASS_COMPL:
        return []DtapIE{}
    case MSG_RR_ASS_FAIL:
        return []DtapIE{}
    case MSG_RR_HANDO_CMD:
        return []DtapIE{}
    case MSG_RR_HANDO_COMPL:
        return []DtapIE{}
    case MSG_RR_HANDO_FAIL:
        return []DtapIE{}
    case MSG_RR_HANDO_INFO:
        return []DtapIE{}
    case MSG_RR_DTM_ASS_CMD:
        return []DtapIE{}

    // Cell Change
    case MSG_RR_CELL_CHG_ORDER:
        return []DtapIE{}
    case MSG_RR_PDCH_ASS_CMD:
        return []DtapIE{}

    // Release
    case MSG_RR_CHAN_REL:
        return []DtapIE{}
    case MSG_RR_PART_REL:
        return []DtapIE{}
    case MSG_RR_PART_REL_COMP:
        return []DtapIE{}

    // Paging & Notification
    case MSG_RR_PAG_REQ_1:
        return []DtapIE{}
    case MSG_RR_PAG_REQ_2:
        return []DtapIE{}
    case MSG_RR_PAG_REQ_3:
        return []DtapIE{}
    case MSG_RR_PAG_RESP:
        return []DtapIE{}
    case MSG_RR_NOTIF_NCH:
        return []DtapIE{}
    case MSG_RR_NOTIF_FACCH:
        return []DtapIE{}
    case MSG_RR_NOTIF_RESP:
        return []DtapIE{}
    case MSG_RR_PACKET_NOTIF:
        return []DtapIE{}

    // Inter-System
    case MSG_RR_UTRAN_CLSM_CHG:
        return []DtapIE{}
    case MSG_RR_CDMA2K_CLSM_CHG:
        return []DtapIE{}
    case MSG_RR_IS_TO_UTRAN_HANDO:
        return []DtapIE{}
    case MSG_RR_IS_TO_CDMA2K_HANDO:
        return []DtapIE{}
    case MSG_RR_GERAN_IU_MODE_CLASSMARK_CHANGE:
        return []DtapIE{}
    case MSG_RR_INTER_SYS_TO_E_UTRAN_HANDO_CMD:
        return []DtapIE{}

    // System Information (BCCH)
    case MSG_RR_SYSINFO_1:
        return []DtapIE{}
    case MSG_RR_SYSINFO_2:
        return []DtapIE{}
    case MSG_RR_SYSINFO_3:
        return []DtapIE{}
    case MSG_RR_SYSINFO_4:
        return []DtapIE{}
    case MSG_RR_SYSINFO_5:
        return []DtapIE{}
    case MSG_RR_SYSINFO_6:
        return []DtapIE{}
    case MSG_RR_SYSINFO_7:
        return []DtapIE{}
    case MSG_RR_SYSINFO_8:
        return []DtapIE{}

    // System Information Extended
    case MSG_RR_SYSINFO_2bis:
        return []DtapIE{}
    case MSG_RR_SYSINFO_2ter:
        return []DtapIE{}
    case MSG_RR_SYSINFO_2quater:
        return []DtapIE{}
    case MSG_RR_SYSINFO_5bis:
        return []DtapIE{}
    case MSG_RR_SYSINFO_5ter:
        return []DtapIE{}
    case MSG_RR_SYSINFO_9:
        return []DtapIE{}
    case MSG_RR_SYSINFO_13:
        return []DtapIE{}
    case MSG_RR_SYSINFO_16:
        return []DtapIE{}
    case MSG_RR_SYSINFO_17:
        return []DtapIE{}
    case MSG_RR_SYSINFO_18:
        return []DtapIE{}
    case MSG_RR_SYSINFO_19:
        return []DtapIE{}
    case MSG_RR_SYSINFO_20:
        return []DtapIE{}

    // Miscellaneous
    case MSG_RR_CHAN_MODE_MODIF:
        return []DtapIE{}
    case MSG_RR_STATUS:
        return []DtapIE{}
    case MSG_RR_CHAN_MODE_MODIF_ACK:
        return []DtapIE{}
    case MSG_RR_FREQ_REDEF:
        return []DtapIE{}
    case MSG_RR_MEAS_REP:
        return []DtapIE{}
    case MSG_RR_CLSM_CHG:
        return []DtapIE{}
    case MSG_RR_CLSM_ENQ:
        return []DtapIE{}
    case MSG_RR_EXT_MEAS_REP:
        return []DtapIE{}
    case MSG_RR_EXT_MEAS_REP_ORD:
        return []DtapIE{}
    case MSG_RR_GPRS_SUSP_REQ:
        return []DtapIE{}
    case MSG_RR_DTM_INFO:
        return []DtapIE{}

    // VGCS
    case MSG_RR_VGCS_UPL_GRANT:
        return []DtapIE{}
    case MSG_RR_UPLINK_RELEASE:
        return []DtapIE{}
    case MSG_RR_UPLINK_FREE:
        return []DtapIE{}
    case MSG_RR_UPLINK_BUSY:
        return []DtapIE{}
    case MSG_RR_TALKER_IND:
        return []DtapIE{}

    // Application Info
    case MSG_RR_APP_INFO:
        return []DtapIE{}
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