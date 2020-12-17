import axios from 'axios'

const baseUrl = '/api/v1'

export const LoginAPI = (username, password) =>
    axios.post(
        baseUrl+'/login', {
            username,
            password
        }
    )

export const RegisterAPI = (username, password) =>
    axios.post(
        baseUrl+'/register', {
            username,
            password
        }
    )