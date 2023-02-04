if (navigator.serviceWorker) {
    navigator.serviceWorker.register('/sw_content.js', { scope: '/' }).then(reg => {
        reg.onupdatefound = () => {
            const installingWorker = reg.installing
            installingWorker.onstatechange = () => {
                if (installingWorker.state === 'installed') {
                    if (navigator.serviceWorker.controller) {
                        console.log('new update available')
                        // new update available
                        window.newUpdateProxy.newUpdate = true
                        caches.keys().then(function (keyList) {
                            return Promise.all(keyList.map(function (key) {
                                return caches.delete(key)
                            })).then(function () {
                                window.setTimeout(
                                    window.location.reload(true),
                                    6000
                                )
                            })
                        })
                    } else {
                        console.log('no update available')
                    }
                }
            }
        }
    }).catch(err => console.error('[SW ERROR]', err))
}
