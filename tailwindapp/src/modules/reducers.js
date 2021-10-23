import { combineReducers } from "redux";

import navbar from "./navbar/reducer";
import profile from "./portfolio/reducer";

export default combineReducers({
    navbar,
    profile
})