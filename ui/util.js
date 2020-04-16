
// Fetch a JSON document.
export function fetchJSON(url, success, errorHandler, method, obj, authToken) {
  if (method === undefined)
    method = 'get'

  const opts = {
    method: method,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
  }

  if (obj !== undefined && obj !== null) {
    opts.body = JSON.stringify(obj)
  }

  if (authToken !== undefined)
    opts.headers.Authorization = 'Bearer ' + authToken

  const pr = fetch(url, opts)

  pr.then(function (response) {
    response.json().then(function (jsonData) {
      if (success)
        success(jsonData)
    })
  })

  pr.catch(function (err) {
    console.log("Error: ", err)
    if(errorHandler)
      errorHandler(err)
  })
}

