import { predictionList } from '$lib/predictions';
import { env } from '$env/dynamic/private';

const SITE_URL = 'https://necrobyte.in'; // Hardcoded for production SEO
const BACKEND_URL = env.PRIVATE_API_URL || 'http://127.0.0.1:8080';

export async function GET() {
    // 1. Static Pages (Home + Prediction Landing Pages)
    const pages = [
        { url: '/', priority: 1.0, changefreq: 'daily' },
        ...predictionList.map(p => ({
            url: `/viral/${p.slug}`,
            priority: 0.8,
            changefreq: 'weekly'
        }))
    ];

    // 2. Dynamic Pages (Recent Predictions) - from Backend
    // Fetching from the backend to get actual generated content specific pages
    try {
        const response = await fetch(`${BACKEND_URL}/api/predictions/recent`);
        if (response.ok) {
            const recentPredictions = await response.json();
            if (Array.isArray(recentPredictions)) {
                recentPredictions.forEach(pred => {
                    // Ensure we have necessary data to construct the URL
                    // Assuming structure: { id: "uuid", page_type: "slug" }
                    if (pred.id && pred.page_type) {
                        pages.push({
                            url: `/viral/${pred.page_type}/${pred.id}`,
                            priority: 0.6,
                            changefreq: 'monthly' // content shouldn't change once generated
                        });
                    }
                });
            }
        }
    } catch (e) {
        console.error('Failed to fetch recent predictions for sitemap', e);
        // Continue generating sitemap with just static pages if backend fails
    }

    // 3. Generate XML
    const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    ${pages
            .map(
                (page) => `
    <url>
        <loc>${SITE_URL}${page.url}</loc>
        <changefreq>${page.changefreq}</changefreq>
        <priority>${page.priority}</priority>
    </url>`
            )
            .join('')}
</urlset>`;

    return new Response(sitemap, {
        headers: {
            'Content-Type': 'application/xml',
            'Cache-Control': 'max-age=3600' // Cache for 1 hour
        }
    });
}
