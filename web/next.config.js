/** @type {import('next').NextConfig} */
const nextConfig = {
  trailingSlash: true,
  basePath: '/loglyzer',
  distDir: 'build',
  async redirects() {
    return [
      {
        source: '/',
        destination: '/analyses',
        permanent: true,
        basePath: false,
      },
      {
        source: '/loglyzer',
        destination: '/loglyzer/analyses',
        permanent: true,
        basePath: false,
      },
    ];
  },
};

module.exports = nextConfig;
