import axios from "axios";

export default async function create(data: any){

    console.log("Intialising creation for: ", data)

    const creation = (await axios.post("http://localhost:8000/items", data)).data

    console.log(creation)

    // return creation;
}