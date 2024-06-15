/** @type {import('next').NextConfig} */
const nextConfig = {
	images: {
		remotePatterns: [
			{
				protocol: "https",
				hostname: "avatars.steamstatic.com",
				port: "",
			}
		]
	}
};

export default nextConfig;
