
/**
 * Fetch a JSON document.
 * 
 * @param {string} url - URL to fetch
 * @param {string} method - 'get', 'post'
 * @param {string} type - 'text' and 'json' are supported
 * @param {Object} obj - request body. This is converted to JSON.
 * @param {string} authToken - JWT authentication token
 * @param {function} success - function(response) called on success
 * @param {function} error - function(message) called on error
 */
export function fetchJSON(url, method, type, obj, authToken, success, error) {
  if (error == undefined) {
    alert(url + ": error handler must be specified")
    return
  }

  if (authToken === undefined) {
    console.log("Error: auth token missing.")
    return
  }

  if (method === undefined)
    method = 'get'

  const opts = {
    method: method,
    headers: {
      'Pragma': 'no-cache',
      'Cache-Control': 'no-cache'
    }
  }

  if (type == 'json') {
    opts.headers["Accept"] = 'application/json'
    opts.headers["Content-Type"] = 'application/json'
  }

  if (obj !== undefined && obj !== null) {
    opts.body = JSON.stringify(obj)
  }

  if (authToken !== undefined)
    opts.headers.Authorization = 'Bearer ' + authToken

  const pr = fetch(url, opts)

  pr.then(function (rs) {
    if (rs.status != 200) {
      rs.json().then(function (msg) {
        if (msg === "")
          msg = rs.statusText
        error(msg, rs.statusText)
      }).catch(function (err) {
        console.log("Error 1: ", err)
        if (error)
          error("Server error.")
      })
    } else {
      rs[type]().then(function (data) {
        //rs.text().then(function (data) {
        if (success)
          success(data)
      })
    }
  })

  pr.catch(function (err) {
    if (error)
      error(err)
  })
}