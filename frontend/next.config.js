/** @type {import('next').NextConfig} */

module.exports = {
  reactStrictMode: true,
  publicRuntimeConfig: {
    apiHostUri: process.env.API_HOST_URI,
  },
};
