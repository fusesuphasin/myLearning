import axios from "axios"

export const TOGGLE_SHOW_PROFILE = 'app/portfolio/TOGGLE_SHOW_PROFILE'
export const SET_SHOW_PROFILE = 'app/portfolio/SET_SHOW_PROFILE'

const LOAD_PRODUCS_REQUEST = 'app/portfolio/LOAD_PRODUCS_REQUEST'
const LOAD_PRODUCS_SUCCESS = 'app/portfolio/LOAD_PRODUCS_SUCCESS'
const LOAD_PRODUCS_FAILURE = 'app/portfolio/LOAD_PRODUCS_FAILURE'

export function toggleShowProfile() {
    return{
        type: TOGGLE_SHOW_PROFILE,
    }
}

export function setShowProfile(showProfile) {
    return{
        type: SET_SHOW_PROFILE,
        payload: {
            showProfile,
        }
    }
}

function loadProducts() {
    return async (dispatch) => {
        dispatch({type: LOAD_PRODUCS_REQUEST})

        try {
            const {data} = await axios.get('http://localhost:4200/purchase/getOrder')

            dispatch({type: LOAD_PRODUCS_SUCCESS, payload: {
                products: data,
            }})
        } catch(err){
            dispatch({type: LOAD_PRODUCS_FAILURE})
        }
    }
   
}

export {
    LOAD_PRODUCS_REQUEST,
    LOAD_PRODUCS_SUCCESS,
    LOAD_PRODUCS_FAILURE,
    loadProducts
}