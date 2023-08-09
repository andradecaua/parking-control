import { useReducer, useState } from "react"
import axios from 'axios'

interface vaga  {
    price: number,
    apelido:  string
    disponivel: boolean
}

interface action {
    payload: any,
    type: string
}

function CriarVaga(){

    const initialVaga:vaga = {price: 0.00, apelido: "", disponivel: true}
    const [vaga, dispatch] = useReducer(reducer, initialVaga)
    const [resultado, setResultado] = useState("")

    function reducer(state: vaga, action:action){
        switch(action.type){
            case "price": 
                return {
                    ...state,
                    price: action.payload
                }
            case "apelido":
                return {
                    ...state,
                    apelido: action.payload
                }
            default: return state 
        }
    }

    function getToken(){
        let config= localStorage.getItem("config")
        let configData = {
            nome: "",
            token: ""
        }
        if (config!= null) {
            configData = JSON.parse(config)
            return configData.token
        }else{
            return "Token Inexistente na configuração"
        }
    }

    async function criarVaga(){
        const response = await axios.post("http://192.168.27.48:80/criar-vaga", vaga, {
            headers: {
                "Content-Type": "application/json",
                "Authorization": getToken()
            }
        })
        let data = await response.data
        setResultado(data.message)
    }

    return(
        <>
            <h2>
                Criar Vaga
            </h2>
            <form action="" method="POST" onSubmit={(event) => {
                event.preventDefault()
                criarVaga()
            }}>
                <label>Apelido da vaga</label>
                <input type="text" name="apelido" onChange={(event) => {dispatch({type: "apelido", payload: event.currentTarget.value})}} />
                <label>Preço da vaga</label>
                <input type="text" name="preço" onChange={(event) => {dispatch({type: "price", payload: Number.parseFloat(event.currentTarget.value)})}}/>
                <button type="submit">Criar Vaga</button>
            </form>
            <span>{resultado}</span>
        </>
    )
}


export default CriarVaga 