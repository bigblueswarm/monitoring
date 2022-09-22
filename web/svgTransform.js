//  This is only a fake transformer to transform svg for tests
module.exports = {
  process () {
    return {
      code: 'module.exports = {};'
    }
  }
}
