import { axiosApi } from "@/api/api.ts";

export interface ArticleCardProps {
    title: string,
    content: string,
    image: string,
    date: string,
}

export const getArticleList = async () => {
    try {
        const response = await axiosApi().get(``)
        return response
    } catch (error: any) {
        return error.response;
    }
}