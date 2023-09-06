import Axios, { AxiosInstance } from 'axios';

export const axiosApi = (): AxiosInstance => {
    const instance = Axios.create({
        baseURL: 'https://superlion.zeabur.app',
        headers: {
            Authorization: ''
        },
    });
    return instance;
}
