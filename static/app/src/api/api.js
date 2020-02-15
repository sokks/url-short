export async function callApi(method, uri, params) {
    let paramList = [];
  
    if (params) {
      for (let name in params) {
        paramList.push(name + '=' + params[name]);
      }
    }
  
    const paramsStr = paramList.join('&');
  
    try {
      const settings = {
          method: method,
          headers: {
              Accept: 'application/json',
          },
          credentials: 'include',
      };
  
      let response = await fetch('/api/' + uri + (paramsStr ? '?' + paramsStr : ''), settings);
  
      if (response.status === 200) {
        let json = await response.json();
        return { json };
      } else {
        return { errors: { message: 'Ошибка ' + response.status }};
      }
    } catch(errors) {
      return { errors };
    }
  }