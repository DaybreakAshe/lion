import { axiosApi } from "../../api/api";

interface Props {
    id:string;
    title:string
}

export const getMyBlogList = async (props: Props) => {
    try {
        const response = await axiosApi().post('/user/posts', props, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })
        return response.data
    } catch (error: any) {
        return error.response;
    }
};

export const getPublicBlogList = async () => {
    try {
        const response = await axiosApi().get('/lion/posts', {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })
        return response.data
    } catch (error: any) {
        return error.response;
    }
};