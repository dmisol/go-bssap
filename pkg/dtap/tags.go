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
        return nil, fmt.Errorf("%w: 0x%02X", ErrProtocolDiscNotExist, pd)
    }
}

func getMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    switch msgType {
    case MSG_MM_IMSI_DETACH_IND:
        return []DtapIE{
            MOBILE_STATION_CLASSMARK_V,
            MOBILE_IDENTITY_LV,
        }, nil

    case MSG_MM_LOC_UPD_ACCEPT:
        return []DtapIE{
            LOCATION_AREA_ID_V,
            MOBILE_IDENTITY_TLV,
            FOLLOW_ON_PROCEED_T,
            CTS_PERMISSION_T,
            EQUIVALENT_PLMNS_TLV,
            EMERGENCY_NUMBER_LIST_TLV,
            PER_MS_T3212_TLV,
            NON_3GPP_NW_PROVIDED_POLICIES_TV,
        }, nil

    case MSG_MM_LOC_UPD_REJECT:
        return []DtapIE{
            REJ_CAUSE_V,
            MM_TIMER_TLV,
        }, nil

    case MSG_MM_LOC_UPD_REQUEST:
        return []DtapIE{
            LOCATION_UPDATING_TYPE_V,
            CIPHERING_KEY_SEQUENCE_NUMBER_V,
            LOCATION_AREA_ID_V,
            MOBILE_STATION_CLASSMARK_V,
            MOBILE_IDENTITY_LV,
            MOBILE_STATION_CLASSMARK_FOR_UMTS_TLV,
            ADDITIONAL_UPDATE_PARAMETERS_TV,
            DEVICE_PROPS_TV,
            MS_NETWORK_FEATURE_SUPPORT_TV,
        }, nil

    case MSG_MM_AUTH_REQ:
        return []DtapIE{
            CIPHERING_KEY_SEQUENCE_NUMBER_V,
            SPARE_HALF_OCTET_V,
            AUTH_PARAM_RAND_V,
            AUTH_PARAM_AUTN_TLV,
        }, nil

    case MSG_MM_AUTH_RESP:
        return []DtapIE{
            AUTH_RESP_PARAM_V,
            AUTH_RESP_PARAM_EXT_TLV,
        }, nil

    case MSG_MM_AUTH_REJ:
        return []DtapIE{}, nil

    case MSG_MM_AUTH_FAIL:
        return []DtapIE{
            REJ_CAUSE_V,
            AUTH_FAILURE_PARAM_TLV,
        }, nil

    case MSG_MM_ID_REQ:
        return []DtapIE{
            IDENTITY_TYPE_V,
            SPARE_HALF_OCTET_V,
        }, nil

    case MSG_MM_ID_RESP:
        return []DtapIE{
            MOBILE_IDENTITY_LV,
            P_TMSI_TYPE_TV,
            ROUTING_AREA_IDENTIFICATION_TLV,
            P_TMSI_SIGNATURE_TLV,
        }, nil

    case MSG_MM_TMSI_REALL_CMD:
        return []DtapIE{
            LOCATION_AREA_ID_V,
            MOBILE_IDENTITY_LV,
        }, nil

    case MSG_MM_TMSI_REALL_COMPL:
        return []DtapIE{}, nil

    case MSG_MM_CM_SERV_ACC:
        return []DtapIE{}, nil

    case MSG_MM_CM_SERV_REJ:
        return []DtapIE{
            REJ_CAUSE_V,
            MM_TIMER_TLV,
        }, nil

    case MSG_MM_CM_SERV_ABORT:
        return []DtapIE{}, nil

    case MSG_MM_CM_SERV_REQ:
        return []DtapIE{
            CM_SERVICE_TYPE_V,
            CIPHERING_KEY_SEQUENCE_NUMBER_V,
            MOBILE_STATION_CLASSMARK_LV,
            MOBILE_IDENTITY_LV,
            PRIORITY_TV,
            ADDITIONAL_UPDATE_PARAMETERS_TV,
            DEVICE_PROPS_TV,
        }, nil

    case MSG_MM_CM_SERV_PROMPT:
        return []DtapIE{PD_AND_SAPI_V}, nil

    case MSG_MM_CM_REEST_REQ:
        return []DtapIE{
            CIPHERING_KEY_SEQUENCE_NUMBER_V,
            SPARE_HALF_OCTET_V,
            MOBILE_STATION_CLASSMARK_LV,
            MOBILE_IDENTITY_LV,
            LOCATION_AREA_ID_TV,
            DEVICE_PROPS_TV,
        }, nil

    case MSG_MM_ABORT:
        return []DtapIE{REJ_CAUSE_V}, nil

    case MSG_MM_NULL:
        return []DtapIE{}, nil

    case MSG_MM_STATUS:
        return []DtapIE{REJ_CAUSE_V}, nil

    case MSG_MM_INFO:
        return []DtapIE{
            FULL_NAME_FOR_NETWORK_TLV,
            SHORT_NAME_FOR_NETWORK_TLV,
            LOCAL_TIME_ZONE_TV,
            UNIVERSAL_TIME_AND_LOCAL_TIME_ZONE_TV,
            LSA_IDENTITY_TLV,
            NETWORK_DAYLIGHT_SAVING_TIME_TLV,
        }, nil

    default:
        return nil, fmt.Errorf("%w: 0x%02X", ErrMsgTypeNotExist, msgType)
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
        return nil, fmt.Errorf("%w: 0x%02X", ErrUnsupportedMsgType, msgType)
    }
}

//ToDO
func getBCCHTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: BCCH protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getCCTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: CC/SS protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getSMSTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: SMS protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getGPRSMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: GPRS MM protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getGPRSSMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: GPRS SM protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getLOCTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: LOC protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getGroupCCTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: Group CC protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getEPSSMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: EPS SM protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getEPSMMMTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: EPS MM protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getGTTTTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: GTTP protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getNCSSTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: SS NC protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getExtendTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: EXTEND protocol not implemented yet", ErrUnsupportedProtocolDisc)
}

func getTestTagsOrder(msgType Msg_Type) ([]DtapIE, error) {
    return nil, fmt.Errorf("%w: TEST protocol not implemented yet", ErrUnsupportedProtocolDisc)
}