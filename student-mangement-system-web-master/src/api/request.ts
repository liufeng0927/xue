import axios from 'axios'
import {useUserStore} from "../store/modules/user";

const service = axios.create({
    baseURL: import.meta.env.VITE_APP_BASE_API,
    timeout: 3000000,
    // 跨域时候允许携带凭证
     //withCredentials: true,
     headers: {
        "Content-Type": "application/json",
     }
})
service.interceptors.request.use(
    config => {
        let userStore = useUserStore();
        let datas = {
            'Content-type': 'application/x-www-form-urlencoded',
            'Authorization': 'Bearer ' + userStore.token
        }
        // @ts-ignore
        config.headers.Authorization = 'Bearer ' + userStore.token
        config.method === 'get' ? config.headers = datas : ''
        return config;
    }
)

export default service
