import type { Category } from "./category";
import type { Comment } from "./comment";
import type { SummitLog } from "./summit_log";
import type { Waypoint } from "./waypoint";

class Trail {
    id?: string;
    name: string;
    location?: string;
    public: boolean;
    distance?: number;
    elevation_gain?: number;
    duration?: number;
    difficulty?: "easy" | "moderate" | "difficult"
    lat?: number;
    lon?: number;
    thumbnail: number;
    photos: string[];
    gpx?: string;
    created?: string;
    category?: string;
    waypoints: string[];
    summit_logs: string[];
    expand: {
        category?: Category;
        waypoints: Waypoint[]
        summit_logs: SummitLog[]
        comments_via_trail: Comment[]
        gpx_data?: string
    }
    tags?: string[];
    description?: string;
    author?: string;

    constructor(name: string,
        params?: {
            id?: string,
            location?: string,
            public?: boolean,
            distance?: number,
            elevation_gain?: number,
            duration?: number,
            difficulty?: "easy" | "moderate" | "difficult",
            lat?: number,
            lon?: number,
            thumbnail?: number,
            photos?: string[],
            gpx?: string,
            category?: Category,
            waypoints?: Waypoint[],
            summit_logs?: SummitLog[],
            comments?: Comment[],
            tags?: string[],
            description?: string
            created?: string
        }

    ) {
        this.id = params?.id;
        this.name = name;
        this.location = params?.location;
        this.public = params?.public ?? false
        this.distance = params?.distance;
        this.elevation_gain = params?.elevation_gain;
        this.duration = params?.duration;
        this.difficulty = params?.difficulty ?? "easy";
        this.lat = params?.lat;
        this.lon = params?.lon;
        this.thumbnail = params?.thumbnail ?? 0;
        this.photos = params?.photos ?? [];
        this.waypoints = [];
        this.summit_logs = [];
        this.gpx = params?.gpx;
        this.expand = {
            category: params?.category,
            waypoints: params?.waypoints ?? [],
            summit_logs: params?.summit_logs ?? [],
            comments_via_trail: params?.comments ?? []
        }
        this.tags = params?.tags ?? []
        this.description = params?.description ?? "";
        this.created = params?.created;
    }
}

interface TrailFilter {
    q: string,
    category: string[],
    difficulty: ("easy" | "moderate" | "difficult")[]
    near: {
        lat?: number,
        lon?: number,
        radius: number
    }
    distanceMin: number,
    distanceMax: number,
    distanceLimit: number,
    elevationGainMin: number;
    elevationGainMax: number;
    elevationGainLimit: number;
    completed?: boolean;
    sort: "name" | "distance" | "elevation_gain" | "created";
    sortOrder: "+" | "-"
}

interface TrailFilterValues {
    min_distance: number,
    max_distance: number,
    min_elevation_gain: number,
    max_elevation_gain: number,
    min_durtation: number,
    max_duration: number,
}

export { Trail };

export type { TrailFilter, TrailFilterValues };