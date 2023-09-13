import Axios, { AxiosInstance } from 'axios';
import { getStoredValue } from '../utils/storage';

export const axiosApi = (): AxiosInstance => {
    const instance = Axios.create({
        baseURL: process.env.REACT_APP_API_URL,
        headers: {
            Authorization: `${getStoredValue('access_token')}` || ''
        },
    });
    return instance;
}
