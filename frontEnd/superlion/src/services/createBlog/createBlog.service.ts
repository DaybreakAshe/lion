import { axiosApi } from "../../api/api";

interface fileProps {
    picture: File,
    busiType: string,
}

export const uploadFile = async (props: fileProps) => {
    try {
        const response = await axiosApi().post('/auth/upload', props, {
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })
        return response.data
    } catch (error: any) {
        return error.response;
    }
}

interface blogProps {
    title: string,
    content: string,
    isPublic: boolean,
    tag: string[],
    category: string[],
}

export const createBlog = async (props: blogProps) => {
    try {
        const response = await axiosApi().post('/auth/createBlog', props)
        return response.data
    } catch (error: any) {
        return error.response;
    }
}