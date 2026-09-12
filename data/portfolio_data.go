package data

import "portfolio-backend/models"

// GetProfile returns the personal overview and contact information for Sharat.
func GetProfile() models.Profile {
	return models.Profile{
		Name:     "Sharat R Kubasad",
		Headline: "Engineer & Cybersecurity Enthusiast",
		Tagline:  "Engineering secure solutions for a safer digital future.",
		Quote:    "The best way to predict the future is to create it",
		Bio: []string{
			"I am an engineering student at SDMCET College in Dharwad, pursuing my passion for technology with a focus on cybersecurity, ethical hacking, bug bounty, and networking. I am deeply interested in exploring the digital world, understanding its intricacies, and ensuring its security.",
			"My journey in the field of cybersecurity began with a curiosity about how systems work and a drive to protect them from malicious threats. As I continue my education, I am dedicated to gaining hands-on experience in ethical hacking and networking, with the ultimate goal of becoming an expert in cybersecurity and making a significant impact in the tech industry.",
			"Outside of academics, I am actively involved in various projects and research, constantly learning new techniques and staying updated with the latest trends in cybersecurity. Whether it's identifying vulnerabilities, securing networks, or participating in bug bounty programs, I am committed to contributing to a safer digital world.",
		},
		Phone:    "+91 7022175067",
		Location: "Dharwad, Karnataka, India",
		SocialLinks: models.SocialLinks{
			GitHub:    "https://github.com/SHARAT1535",
			LinkedIn:  "https://www.linkedin.com/in/sharat-kubasad/",
			Instagram: "https://www.instagram.com/sharat__kubasad/",
			Facebook:  "https://www.facebook.com/sharat.r.kubasad/",
			Twitter:   "https://twitter.com/sharatkubasad",
		},
	}
}

// GetSkills returns the list of technical skills and proficiency levels.
func GetSkills() []models.Skill {
	return []models.Skill{
		{Name: "Go (Golang)", Proficiency: 75, Status: "Proficient"},
		{Name: "Networking", Proficiency: 80, Status: "Advanced"},
		{Name: "Linux", Proficiency: 80, Status: "Advanced"},
		{Name: "Python", Proficiency: 70, Status: "Proficient"},
		{Name: "C/C++", Proficiency: 70, Status: "Proficient"},
		{Name: "Cybersecurity & Bug Bounty", Proficiency: 75, Status: "Active"},
		{Name: "App Development", Proficiency: 65, Status: "Ongoing"},
	}
}

// GetProjects returns all portfolio projects with full descriptions and metadata.
func GetProjects() []models.Project {
	return []models.Project{
		{
			ID:      "news-app",
			Title:   "News App",
			Summary: "A dynamic web application developed using Flask, designed to provide users with the latest headlines across multiple categories using NewsAPI.",
			KeyFeatures: []string{
				"Dynamic Content: Real-time news article ingestion across General, Sports, Tech, and Business.",
				"Responsive Design: Styled with clean CSS and responsive layout across mobile and desktop.",
				"Category Filtering: Fast categorized navigation for streamlined news browsing.",
				"External API Integration: Clean integration with NewsAPI for fresh data.",
			},
			Technologies: []string{"Flask", "Python", "NewsAPI", "HTML5", "CSS3"},
			Goals:        "Demonstrate proficiency in web application routing, external API consumption, and responsive interface design.",
		},
		{
			ID:      "packet-sniffer",
			Title:   "Network Packet Sniffer",
			Summary: "A network analysis script implemented in Python using the Scapy library to capture and inspect HTTP traffic in real time.",
			KeyFeatures: []string{
				"Interface Selection: Allows selecting specific network interfaces via CLI flags.",
				"Real-Time Capture: Captures and dissects packets on the fly.",
				"HTTP Payload Inspection: Extracts host, request path, and headers.",
				"Sensitive Data Detection: Scans payloads for credentials, tokens, and sensitive strings.",
			},
			Technologies: []string{"Python", "Scapy", "Networking Protocols", "Linux"},
			Goals:        "Monitor network traffic, understand protocol structures, and analyze payload security.",
		},
		{
			ID:      "temple-connect",
			Title:   "Temple Connect Android App",
			Summary: "An Android application designed to bridge the gap between devotees and temples, offering real-time event updates and notifications.",
			KeyFeatures: []string{
				"Temple Directory: Curated list of temples with history, timings, and services.",
				"Nearest Temple Finder: Uses GPS location to find temples in proximity.",
				"Donation & Support: Integrated donation module for devotees.",
				"Real-time Updates: Push notifications for temple festivals and events.",
			},
			Technologies: []string{"Android Studio", "Java", "Kotlin", "Firebase"},
			Goals:        "Provide a unified, intuitive mobile experience connecting communities with local cultural institutions.",
		},
		{
			ID:      "coffee-machine",
			Title:   "Automatic Coffee Machine",
			Summary: "An IoT hardware project utilizing Arduino microcontroller, servo motors, and relays to automate coffee preparation.",
			KeyFeatures: []string{
				"Servo Motor Dispensing: Precise angle control for ingredient measurement.",
				"Sensor Driven: Triggered via manual push-button or infrared proximity sensor.",
				"Relay Pump Control: Automated liquid dispensing with safety cutoffs.",
				"Serial Diagnostics: Real-time status output over USB serial.",
			},
			Technologies: []string{"Arduino", "C/C++", "Sensors", "Hardware Interfacing"},
			Goals:        "Automate physical beverage preparation with microcontrollers and embedded sensors.",
		},
		{
			ID:      "portfolio-site",
			Title:   "Personal Developer Portfolio & API",
			Summary: "A modern developer portfolio powered by a Go REST backend and React frontend, migrating from the original Flask implementation.",
			KeyFeatures: []string{
				"Go REST API: Ultra-fast compiled API endpoints with zero bloat.",
				"Contact Dispatcher: Secure SMTP integration with validation.",
				"Modern Architecture: Decoupled frontend on Vercel and backend API.",
			},
			Technologies: []string{"Go (Golang)", "React", "REST API", "Docker", "Vercel"},
			Goals:        "Showcase full-stack engineering skills, modern Go backend practices, and clean architecture.",
		},
	}
}

// GetBlogs returns published technical articles and walkthroughs.
func GetBlogs() []models.BlogPost {
	return []models.BlogPost{
		{
			Title:       "Networking",
			URL:         "https://medium.com/@wandesrtech/computer-networking-dfa73e7ebc74",
			Description: "A detailed look into computer networking concepts, architectures, and techniques.",
		},
		{
			Title:       "Network Devices",
			URL:         "https://medium.com/@wandesrtech/network-devices-e94c83e56256",
			Description: "Exploring various network devices (routers, switches, hubs) and their functionalities.",
		},
		{
			Title:       "Transmission Medium",
			URL:         "https://medium.com/@wandesrtech/transmission-medium-f12edd6592c1",
			Description: "Understanding guided and unguided transmission media used in telecommunications.",
		},
		{
			Title:       "OSI Reference Model",
			URL:         "https://medium.com/@wandesrtech/osi-reference-model-61e4e72c6546",
			Description: "An overview of the OSI 7-layer reference model and how each layer operates.",
		},
		{
			Title:       "Kioptrix Level 1 Walkthrough",
			URL:         "https://medium.com/@wandesrtech/kioptrix-level-1-walkthrough-277f460faf87",
			Description: "A hands-on walkthrough exploiting the Kioptrix Level 1 vulnerable virtual machine.",
		},
		{
			Title:       "Kioptrix Level 2 Walkthrough",
			URL:         "https://medium.com/@wandesrtech/kioptrix-level-2-82150fe1c200",
			Description: "Detailed step-by-step penetration testing guide for solving Kioptrix Level 2.",
		},
		{
			Title:       "Metasploitable 2 Walkthrough",
			URL:         "https://medium.com/@wandesrtech/metasploitable-2-walkthrough-c2cd662c9e76",
			Description: "Step-by-step walkthrough of Metasploitable 2 vulnerability assessment and exploitation.",
		},
	}
}
