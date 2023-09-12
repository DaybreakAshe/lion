import Axios, { AxiosInstance } from 'axios';
import { getStoredValue } from '../utils/storage';

export const axiosApi = (): AxiosInstance => {
    const instance = Axios.create({
        baseURL: 'https://superlion.zeabur.app/lion',
        headers: {
            Authorization: `${getStoredValue('access_token')}` || ''
        },
    });
    return instance;
}
