import { axiosApi } from "../../api/api";


export const getUserInfo = async (params:any) => {
    try {
        const response = await axiosApi().post('/lion/login', params)
        return response.data
    } catch (error: any) {
        return error.response;
    }
}