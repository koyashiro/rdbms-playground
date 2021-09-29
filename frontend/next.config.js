/** @type {import('next').NextConfig} */

const validateNotEmptyEnv = (source) => {
  if (source != null && source !== "") {
    return source;
  } else {
    throw new Error("Environment variable validation failed");
  }
};

module.exports = {
  reactStrictMode: true,
  publicRuntimeConfig: {
    apiHostUri: validateNotEmptyEnv(process.env.API_HOST_URI),
  },
};
