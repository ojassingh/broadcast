import axios from "axios";

export default async function readOne(data: any){

    const id = data.id;
    const read = axios.get(`http://localhost:8000/items/${id}`, data)

    return read;
}

export async function readAll(){
    const read = axios.get(`http://localhost:8000/items`)
    return read;
}

