export default {
    /**
     *
     * @param {any} value
     */
    isNumber: function (value) {
      return !isNaN(parseFloat(value)) && isFinite(value);
    },
  };
  