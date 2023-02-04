import { useContext } from 'react';
import { UserContext } from './components/UserContext';

export const useAPI = () => {
    const { userStatus, setUserStatus } = useContext(UserContext);

    const dev = true;
    const basePath = dev ? "http://127.0.0.1:9000/" : "/"

    const api = (path, method = "GET", body = {}) => {
        const base_path = basePath + "api/"

        return new Promise((resolve) => {
            const h = { "Content-Type": "application/json" }
            if (userStatus.connected) {
                h["Authorization"] = "Bearer " + userStatus.token
            }
            const init = {
                method: method,
                headers: h
            }
            if (method !== "GET" && body !== {}) {
                init.body = JSON.stringify(body)
            }
            fetch(base_path + path, init).then((resp) => {
                resp.json().then((data) => {
                    if (resp.status !== 200) {
                        if (resp.status === 401) {
                            // token not recognised by the server
                            if (path !== "user/login") {
                                userStatus.connected = false
                                userStatus.error = false
                                setUserStatus(userStatus)
                                window.location.replace(window.location.protocol + "//" + window.location.host)
                            }
                        } else {
                            error(resp.status + " - " + data.data)
                        }
                        resolve(null)
                    }
                    userStatus.error = false
                    setUserStatus(userStatus)
                    resolve(data.data)
                }).catch((err) => {
                    console.error(err)
                    resolve(null)
                })
            }).catch((err) => {
                error(err)
                resolve(null)
            })
        })
    }

    function error(err) {
        console.error("request error: " + err)
        userStatus.error = true
        userStatus.errorMessage = err
        setUserStatus(userStatus)
    }

    return { api, basePath };
}