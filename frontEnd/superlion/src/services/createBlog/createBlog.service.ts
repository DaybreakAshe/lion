import { axiosApi } from "../../api/api";

interface fileProps {
    picture: any,
    busiType: string,
}
export const uploadFile = async (prop: fileProps) => {
    try {
        const response = await axiosApi().post('/auth/upload', prop)
        return response.data
    } catch (error: any) {
        return error.response;
    }
}