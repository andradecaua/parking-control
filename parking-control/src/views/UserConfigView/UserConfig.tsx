import { useReducer } from "react"
import "./userConfig.scss"

interface configDataInteface {
    nome: string,
    token: string
}

interface actionInterface {
        payload: string,
        type: string
}

function UserConfig(){

    const initialState:configDataInteface = {
        nome: "",
        token: ""
    }
    const [configData, dispatch] = useReducer(reducer, initialState)

    function saveConfig(){
        let getActualConfig = localStorage.getItem('config')
        if(getActualConfig == null){
            localStorage.setItem('config', JSON.stringify(configData))
        }else{
            localStorage.removeItem('config')
            localStorage.setItem('config', JSON.stringify(configData))
        }
    }

    function reducer(state:configDataInteface, action: actionInterface){
        switch(action.type){
            case "nome":
               return {
                    ...state,
                    nome: action.payload
                }
            case "token": 
                return {
                    ...state,
                    token: action.payload
                }
            default: return state
        }
    }

    return(
        <main>
            <form action="javascript:void(0)" onSubmit={saveConfig}>
                <label htmlFor="nome">
                    Nome
                </label>
                <input name="nome" type="text" onChange={(event) => dispatch({type: "nome", payload: event.currentTarget.value})} />
                <label>
                    Token de Acesso
                </label>
                <input name="token" type="password" onChange={(event) => dispatch({type: "token", payload: event.currentTarget.value})} />
                <button type="submit">Salvar Configurações</button>
            </form>
        </main>
    )
}

export default UserConfig