package dtap

import (
    "fmt"
)

func GetTagsOrder(pd PD_Type, msgType Msg_Type) ([]DtapIE, error) {
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
        return nil, fmt.Errorf("unsupported protocol discriminator: 0x%02X", pd)
    }
}

func getMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    switch msgType {
    case MSG_MM_IMSI_DETACH_IND:
        return []DtapIE{
            MS_CLASSMARK_1,
            M_IDENTITY_1,
        }, nil

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
        }, nil

    case MSG_MM_LOC_UPD_REJECT:
        return []DtapIE{
            REJ_CAUSE,
            MM_TIMER,
        }, nil

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
        }, nil

    case MSG_MM_AUTH_REQ:
        return []DtapIE{
            CIPH_KEY_SEQ_NUM,
            SPARE_HALF_OCT,
            AUTH_PARAM_RAND,
            AUTH_PARAM_AUTN,
        }, nil

    case MSG_MM_AUTH_RESP:
        return []DtapIE{
            AUTH_RESP_PARAM,
            AUTH_RESP_PARAM_EXT,
        }, nil

    case MSG_MM_AUTH_REJ:
        return []DtapIE{}, nil

    case MSG_MM_AUTH_FAIL:
        return []DtapIE{
            REJ_CAUSE,
            AUTH_FAIL_PARAM,
        }, nil

    case MSG_MM_ID_REQ:
        return []DtapIE{
            ID_TYPE,
            SPARE_HALF_OCT,
        }, nil

    case MSG_MM_ID_RESP:
        return []DtapIE{
            M_IDENTITY_1,
            P_TMSI_TYPE,
            ROUT_AREA_ID_2,
            P_TMSI_SIGN_2,
        }, nil

    case MSG_MM_TMSI_REALL_CMD:
        return []DtapIE{
            LOC_AREA_ID,
            M_IDENTITY_1,
        }, nil

    case MSG_MM_TMSI_REALL_COMPL:
        return []DtapIE{}, nil

    case MSG_MM_CM_SERV_ACC:
        return []DtapIE{}, nil

    case MSG_MM_CM_SERV_REJ:
        return []DtapIE{
            REJ_CAUSE,
            MM_TIMER,
        }, nil

    case MSG_MM_CM_SERV_ABORT:
        return []DtapIE{}, nil

    case MSG_MM_CM_SERV_REQ:
        return []DtapIE{
            CM_SERVICE_TYPE,
            CIPH_KEY_SEQ_NUM,
            MS_CLASSMARK_2,
            M_IDENTITY_1,
            PRIOR_LVL,
            ADD_UPD_PARAMS,
            DEVICE_PROPS,
        }, nil

    case MSG_MM_CM_SERV_PROMPT:
        return []DtapIE{PD_AND_SAPI}, nil

    case MSG_MM_CM_REEST_REQ:
        return []DtapIE{
            CIPH_KEY_SEQ_NUM,
            SPARE_HALF_OCT,
            MS_CLASSMARK_2,
            M_IDENTITY_1,
            LOC_AREA_ID,
            DEVICE_PROPS,
        }, nil

    case MSG_MM_ABORT:
        return []DtapIE{REJ_CAUSE}, nil

    case MSG_MM_NULL:
        return []DtapIE{}, nil

    case MSG_MM_STATUS:
        return []DtapIE{REJ_CAUSE}, nil

    case MSG_MM_INFO:
        return []DtapIE{
            FNAME_F_NET,
            SNAME_F_NET,
            TIME_ZONE,
            TIME_ZONE_AND_TIME,
            LSA_IDEN,
            DAY_SAVING_TIME,
        }, nil

    default:
        return nil, fmt.Errorf("unsupported MM message type: 0x%02X", msgType)
    }
}

func getRRTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
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
        }, nil

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
        }, nil

    case MSG_RR_CIPH_M_CMD:
        return []DtapIE{
            CIPHERING_MODE_SETTING_V,
            CIPHER_RESPONSE_V,
        }, nil

    case MSG_RR_CIPH_M_COMPL:
        return []DtapIE{
            ME_IDENTITY_TLV,
        }, nil

    case MSG_RR_ASS_COMPL:
        return []DtapIE{
            RR_CAUSE_V,
        }, nil

    case MSG_RR_ASS_FAIL:
        return []DtapIE{
            RR_CAUSE_V,
        }, nil

    case MSG_RR_ASS_CMD:
        return []DtapIE{
            DESCRIPTION_OF_FIRST_CHANNEL_AFTER_TIME_V,
            POWER_COMMAND_V,
            FREQUENCY_LIST_AFTER_TIME_TLV,
            CELL_CHANNEL_DESCRIPTION_TV,
            DESCRIPTION_OF_MULTISLOT_CONFIGURATION_TLV,
            MODE_OF_FIRST_CHANNEL_CHANNEL_SET_1_TV,
            MODE_OF_CHANNEL_SET_2_TV,
            MODE_OF_CHANNEL_SET_3_TV,
            MODE_OF_CHANNEL_SET_4_TV,
            MODE_OF_CHANNEL_SET_5_TV,
            MODE_OF_CHANNEL_SET_6_TV,
            MODE_OF_CHANNEL_SET_7_TV,
            MODE_OF_CHANNEL_SET_8_TV,
            DESCRIPTION_OF_SECOND_CHANNEL_AFTER_TIME_TV,
            MODE_OF_SECOND_CHANNEL_TV,
            MOBILE_ALLOCATION_AFTER_TIME_TLV,
            STARTING_TIME_TV,
            FREQUENCY_LIST_BEFORE_TIME_TLV,
            DESCRIPTION_OF_FIRST_CHANNEL_BEFORE_TIME_TV,
            DESCRIPTION_OF_SECOND_CHANNEL_BEFORE_TIME_TV,
            FREQUENCY_CHANNEL_SEQUENCE_BEFORE_TIME_TV,
            MOBILE_ALLOCATION_BEFORE_TIME_TLV,
            CIPHER_MODE_SETTING_TV,
            VGCS_TARGET_MODE_INDICATION_TLV,
            MULTI_RATE_CONFIGURATION_TLV,
            VGCS_CIPHERING_PARAMETERS_TLV,
            EXTENDED_TSC_SET_TV,
            EXTENDED_TSC_SET_BEFORE_TIME_TV,
        }, nil

    case MSG_RR_HANDO_CMD:
        return []DtapIE{
            CELL_DESCRIPTION_V,
            DESCRIPTION_OF_FIRST_CHANNEL_AFTER_TIME_V,
            HANDOVER_REFERENCE_V,
            POWER_COMMAND_AND_ACCESS_TYPE_V,
            SYNCHRONIZATION_INDICATION_TV,
            FREQUENCY_SHORT_LIST_AFTER_TIME_TV,
            FREQUENCY_LIST_AFTER_TIME_TLV,
            CELL_CHANNEL_DESCRIPTION_TV,
            DESCRIPTION_OF_MULTISLOT_CONFIGURATION_TLV,
            MODE_OF_FIRST_CHANNEL_CHANNEL_SET_1_TV,
            MODE_OF_CHANNEL_SET_2_TV,
            MODE_OF_CHANNEL_SET_3_TV,
            MODE_OF_CHANNEL_SET_4_TV,
            MODE_OF_CHANNEL_SET_5_TV,
            MODE_OF_CHANNEL_SET_6_TV,
            MODE_OF_CHANNEL_SET_7_TV,
            MODE_OF_CHANNEL_SET_8_TV,
            DESCRIPTION_OF_SECOND_CHANNEL_AFTER_TIME_TV,
            MODE_OF_SECOND_CHANNEL_TV,
            FREQUENCY_CHANNEL_SEQUENCE_AFTER_TIME_TV,
            MOBILE_ALLOCATION_AFTER_TIME_TLV,
            STARTING_TIME_TV,
            REAL_TIME_DIFFERENCE_TLV,
            TIMING_ADVANCE_TV,
            FREQUENCY_SHORT_LIST_BEFORE_TIME_TV,
            FREQUENCY_LIST_BEFORE_TIME_TLV,
            DESCRIPTION_OF_FIRST_CHANNEL_BEFORE_TIME_TV,
            DESCRIPTION_OF_SECOND_CHANNEL_BEFORE_TIME_TV,
            FREQUENCY_CHANNEL_SEQUENCE_BEFORE_TIME_TV,
            MOBILE_ALLOCATION_BEFORE_TIME_TLV,
            CIPHER_MODE_SETTING_TV,
            VGCS_TARGET_MODE_INDICATION_TLV,
            MULTI_RATE_CONFIGURATION_TLV,
            DYNAMIC_ARFCN_MAPPING_TLV,
            VGCS_CIPHERING_PARAMETERS_TLV,
            DEDICATED_SERVICE_INFORMATION_TV,
            PLMN_INDEX_TV,
            EXTENDED_TSC_SET_TV,
            EXTENDED_TSC_SET_BEFORE_TIME_TV,
        }, nil

    case MSG_RR_HANDO_COMPL:
        return []DtapIE{
            RR_CAUSE_V,
            MOBILE_OBSERVED_TIME_DIFF_TLV,
            MOBILE_OBSERVED_TIME_DIFF_ON_HYPERFRAME_LVL_TLV,
        }, nil

    case MSG_RR_HANDO_FAIL:
        return []DtapIE{
            RR_CAUSE_V,
            PS_CAUSE_TV,
        }, nil

    case MSG_RR_CHAN_REL:
        return []DtapIE{
            RR_CAUSE_V,
            BA_RANGE_TLV,
            GROUP_CHANNEL_DESCRIPTION_TLV,
            GROUP_CIPHER_KEY_NUMBER_TV,
            GPRS_RESUMPTION_TV,
            BA_LIST_PREF_TLV,
            UTRAN_FREQ_LIST_TLV,
            CELL_CHANNEL_DESCRIPTION_TV,
            CELL_SELECTION_INDICATOR_AFTER_RELEASE_TLV,
            ENHANCED_DTM_CS_RELEASE_INDICATION_TV,
            VGCS_CIPHERING_PARAMETERS_TLV,
            GROUP_CHANNEL_DESCRIPTION_2_TLV,
            TALKER_IDENTITY_TLV,
            TALKER_PRIORITY_STATUS_TLV,
            VGCS_AMR_CONFIGURATION_TLV,
            INDIVIDUAL_PRIORITIES_TLV,
        }, nil

    case MSG_RR_PAG_REQ_1:
        return []DtapIE{
            PAGE_MODE_V,
            CHANNELS_NEEDED_FOR_MOBILES_1_AND_2_V,
            MOBILE_IDENTITY_1_LV,
            MOBILE_IDENTITY_2_TLV,
            //P1_REST_OCTETS_V, не поддерживается еще
        }, nil

    case MSG_RR_PAG_REQ_2:
        return []DtapIE{
            PAGE_MODE_V,
            CHANNELS_NEEDED_FOR_MOBILES_1_AND_2_V,
            MOBILE_IDENTITY_1_V,
            MOBILE_IDENTITY_2_V,
            MOBILE_IDENTITY_3_TLV,
            //P2_REST_OCTETS_V, не поддерживается еще
        }, nil

    case MSG_RR_PAG_REQ_3:
        return []DtapIE{
            PAGE_MODE_V,
            CHANNELS_NEEDED_FOR_MOBILES_1_AND_2_V,
            MOBILE_IDENTITY_1_V,
            MOBILE_IDENTITY_2_V,
            MOBILE_IDENTITY_3_V,
            MOBILE_IDENTITY_4_V,
            P3_REST_OCTETS_V,
        }, nil

    case MSG_RR_PAG_RESP:
        return []DtapIE{
            CIPHERING_KEY_SEQUENCE_NUMBER_V,
            SPARE_HALF_OCTET_V,
            MOBILE_STATION_CLASSMARK_LV,
            MOBILE_IDENTITY_LV,
            ADDITIONAL_UPDATE_PARAMETERS_TV,
        }, nil

    default:
        return nil, fmt.Errorf("unsupported RR message type: 0x%02X", msgType)
    }
}

//ToDO
func getBCCHTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("BCCH not implemented yet for msg type: 0x%02X", msgType)
}

func getCCTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("CC not implemented yet for msg type: 0x%02X", msgType)
}

func getSMSTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("SMS not implemented yet for msg type: 0x%02X", msgType)
}

func getGPRSMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("GPRS MM not implemented yet for msg type: 0x%02X", msgType)
}

func getGPRSSMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("GPRS SM not implemented yet for msg type: 0x%02X", msgType)
}

func getLOCTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("LOC not implemented yet for msg type: 0x%02X", msgType)
}

func getGroupCCTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("Group CC not implemented yet for msg type: 0x%02X", msgType)
}

func getEPSSMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("EPS SM not implemented yet for msg type: 0x%02X", msgType)
}

func getEPSMMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("EPS MM not implemented yet for msg type: 0x%02X", msgType)
}

func getGTTTTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("GTTP not implemented yet for msg type: 0x%02X", msgType)
}

func getNCSSTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("SS NC not implemented yet for msg type: 0x%02X", msgType)
}

func getExtendTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("EXTEND not implemented yet for msg type: 0x%02X", msgType)
}

func getTestTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("TEST not implemented yet for msg type: 0x%02X", msgType)
}