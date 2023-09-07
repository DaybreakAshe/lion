import { axiosApi } from "../../api/api";
export interface ArticleCardProps {
    title: string,
    content: string,
    image: string,
    date: string,
}

export const getArticleList = async () => {
    try {
        const response = await axiosApi().get('/hello')
        return response.data
    } catch (error: any) {
        return error.response;
    }
}
