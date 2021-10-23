import React, { useEffect } from 'react'
import "../../../styles/main.css"
import {useSelector, useDispatch} from "react-redux"
import * as actions from '../actions'

const education = [
    {
        img : "../img/3.jpg",
        year : "2554 - 2560",
        school : "SRW",
        level : "Medium School"
    },
    {
        img : "../img/4.png",
        year : "2560 - 2564",
        school : "SUT",
        level : "High School"
    }
]

function Profile() {
    const dispatch = useDispatch()
    const toggleShowProfile = () => {
        dispatch(actions.toggleShowProfile())
        console.log(products)
    }
    const isShowProfile = true;
    const isShow = useSelector(state => state.profile.showProfile) // from main state. in reducers.js
    const { isLoading, items: products} = useSelector(state => state.profile)


    useEffect(() => {
        const action = actions.setShowProfile(isShowProfile)
        dispatch(action)

        const loadProduct = actions.loadProducts()
        dispatch(loadProduct)
    }, [dispatch,isShowProfile])

    return (
      <div class="bg-gray-100 "> 
        <div >
            <button onClick={toggleShowProfile} class="px-4 " /* onClick={()=>setIsShow(!isShow)} */>
                <h1 class="text-4xl font-bold text-red-500">Profile</h1>
                <svg  xmlns="http://www.w3.org/2000/svg" class="px-4 h-20 w-22 fill-current hover:text-yellow-500 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path class="" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
                </svg>
            </button>
        </div>

        {isShow && 
            <div>
                <div class="px-8 py-12 max-w-md mx-auto">
                    <h2 class="text-4xl font-bold text-blue-500 leading-snug">Profile</h2>
                    <img class="h-20 w-25 mt-4 rounded-lg shadow-2xl" src="../img/2.jpg" alt="Workcation"/>
                    <h1 class="text-xl mt-4 font-bold text-left text-blue-500 leading-snug">NAME : 
                        <span class="text-red-500">SUPHASIN YOSANG</span>
                    </h1>
                <p class="mt-3 text-gray-500">I live in sisaket. I graduate in bachelor degree with Suranaree University of Technology in Computer Engineer.</p>
                </div>
                <div class="px-8 pr-12 max-w-md mx-auto">
                    <h2 class="text-yellow-500 text-xl mb-3">Education</h2>
                {
                    education.map((item => {
                        return (
                        
                                <div class="flex items-center shadow-lg ">
                                    <img class="h-20 w-25 rounded-md " src={item.img} alt="" />
                                    <div class="px-10 pr-10">
                                        <h3 class="text-lg text-gray-800 text-right">{item.level}</h3>
                                        <p>{item.year}</p>
                                        <div class="mt-2">
                                            <a href="" class="text-red-400  text-indigo hover:text-blue-900">{item.school}</a>
                                        </div>
                                    </div>
                                </div>       
                        )
                    }))
                }

                <h2 class="text-xl pt-4 pb-3 text-pink-500 leading-snug">Activity</h2>
                
                </div>

                <div class="w-72 t-5 max-w-md mx-auto rounded-md vorder overflow-hidden bg-white shadow-xl">
                    <img class=" h-24 w-full object-cover" src="../img/5.jpg" alt=""/>
                    <div class="px-5 -mt-12">
                        <div class="p-6" >
                            <div class="mt-1 text-gray-900">Position : Photographer</div>
                            <div class="mt-1 text-gray-500">CLUB : Photograph</div>
                        </div>
                    </div>
                
                </div>

            

                <div class="mt-4 pl-10 max-w-md mx-auto">
                    <a href="#" class=" rounded-lg shadow-lg text-white bg-blue-400 px-3 py-2 mr-auto ml-auto inline-block hover:bg-pink-400 hover:text-purple-300 focus:outline-none focus:shadow-outline active:bg-indigo-600 ">FACEBOOK</a>
                    <a href="#" class="rounded-lg shadow-lg text-white bg-blue-400 px-3 py-2 mr-auto ml-auto inline-block hover:bg-pink-400 hover:text-purple-300 focus:outline-none focus:shadow-outline active:bg-indigo-600 ">MAIL</a>
                    <a href="#" class="btn-blue">GITHUB</a>
                    <a href="#" class="btn-blue">PORTFOLIO</a>
                </div>
            </div>
        }
        
      </div>
    )
}

export default Profile

