import axios from 'axios'

const baseUrl = '/api/v1'

export const DeleteNodeAPI = (ip) =>
    axios.post(baseUrl+'/delete_node', {
            ip: ip,
        }
    )

export const AddNodeAPI = (ip, host_name, ssh_username, ssh_password) =>
    axios.post(baseUrl+'/add_node', {
            ip: ip,
            host_name: host_name,
            ssh_username: ssh_username,
            ssh_password: ssh_password
        }
    )