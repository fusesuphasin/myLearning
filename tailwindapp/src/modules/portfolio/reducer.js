import { SET_SHOW_PROFILE, TOGGLE_SHOW_PROFILE } from "./actions"
import { LOAD_PRODUCS_FAILURE, LOAD_PRODUCS_SUCCESS, LOAD_PRODUCS_REQUEST } from "./actions"

const initialState = {
    showProfile: false,
    isLoading: false,
    items: []
}

//ใช้เพียวฟังก์ชันค่าไม่เปลี่ยนแปลง

export default function(state = initialState, action){
    switch(action.type) {
        case LOAD_PRODUCS_REQUEST:
            return {...state, isLoading: true, items: []}
            case LOAD_PRODUCS_SUCCESS:
                return {...state, isLoading: false, items: action.payload.products}
                case LOAD_PRODUCS_FAILURE:
                    return {...state, isLoading: false}
        case TOGGLE_SHOW_PROFILE:
            return {...state, showProfile : !state.showProfile}
        case SET_SHOW_PROFILE:
            return {...state, showProfile : action.payload.showProfile}
        default:
            return state
    }
}

