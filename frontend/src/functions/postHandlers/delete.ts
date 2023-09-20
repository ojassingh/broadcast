import axios from "axios";

export default function del(id: any){

    try{
        console.log("Initiating deletion for id: ", id)
        const deletion = axios.delete(`http://localhost:8000/items/${id}`)
         return deletion;
    }
    catch(err){
        console.log(err)
    }
}