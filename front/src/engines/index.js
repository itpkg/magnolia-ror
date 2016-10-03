import auth from './auth'

const engines = {
  auth
}

export default {
  routes () {
    return Object.keys(engines).reduce(function (obj, en) {
      return obj.concat(engines[en].routes)
    }, [])
  }
}
