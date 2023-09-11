



import { axiosApi } from "../../api/api";

interface tokenProps {

}

export const getUserInfo = async (params:any) => {
    try {
        const response = await axiosApi().post('/login', params)
        return response.data
    } catch (error: any) {
        return error.response;
    }
}