import React from 'react'
import Profile from './modules/portfolio/components/profile'
import { Provider } from "react-redux"
import configureStore from './store/configureStore'


const store = configureStore()

function App() {
  return (
    <Provider store={store}>
        <div>
          <Profile />
        </div>
    </Provider>
   
  )
}

export default App

