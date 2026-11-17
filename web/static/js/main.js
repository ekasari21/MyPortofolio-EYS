// Portfolio JavaScript Enhancements

document.addEventListener('DOMContentLoaded', () => {
    initThemeToggle();
    initProjectFilters();
});

// 1. Theme Toggle (Bootstrap 5.3 Color Modes)
function initThemeToggle() {
    const themeToggleBtn = document.getElementById('themeToggleBtn');
    const themeIcon = document.getElementById('themeIcon');
    const themeText = document.getElementById('themeText');
    const htmlElement = document.documentElement;

    const savedTheme = localStorage.getItem('portfolio_theme') || 
        (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');

    setTheme(savedTheme);

    if (themeToggleBtn) {
        themeToggleBtn.addEventListener('click', () => {
            const currentTheme = htmlElement.getAttribute('data-bs-theme');
            const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
            setTheme(newTheme);
            localStorage.setItem('portfolio_theme', newTheme);
        });
    }

    function setTheme(theme) {
        htmlElement.setAttribute('data-bs-theme', theme);
        if (themeIcon) {
            if (theme === 'dark') {
                themeIcon.className = 'bi bi-sun-fill text-warning';
                if (themeText) themeText.textContent = 'Light';
            } else {
                themeIcon.className = 'bi bi-moon-stars text-secondary';
                if (themeText) themeText.textContent = 'Dark';
            }
        }
    }
}

// 2. Interactive Project Filtering & Search
function initProjectFilters() {
    const filterButtons = document.querySelectorAll('.filter-btn');
    const searchInput = document.getElementById('projectSearchInput');
    const projectItems = document.querySelectorAll('.project-item');
    const noProjectsMsg = document.getElementById('noProjectsMessage');

    if (!filterButtons.length || !projectItems.length) return;

    let activeFilter = 'all';
    let searchQuery = '';

    filterButtons.forEach(btn => {
        btn.addEventListener('click', () => {
            filterButtons.forEach(b => {
                b.classList.remove('btn-primary', 'active');
                b.classList.add('btn-outline-secondary');
            });
            btn.classList.remove('btn-outline-secondary');
            btn.classList.add('btn-primary', 'active');

            activeFilter = btn.getAttribute('data-filter');
            applyFilters();
        });
    });

    if (searchInput) {
        searchInput.addEventListener('input', (e) => {
            searchQuery = e.target.value.toLowerCase().trim();
            applyFilters();
        });
    }

    function applyFilters() {
        let visibleCount = 0;

        projectItems.forEach(item => {
            const category = item.getAttribute('data-category');
            const name = (item.getAttribute('data-name') || '').toLowerCase();
            const tech = (item.getAttribute('data-tech') || '').toLowerCase();

            const matchesCategory = (activeFilter === 'all' || category === activeFilter);
            const matchesSearch = !searchQuery || name.includes(searchQuery) || tech.includes(searchQuery);

            if (matchesCategory && matchesSearch) {
                item.classList.remove('d-none');
                visibleCount++;
            } else {
                item.classList.add('d-none');
            }
        });

        if (noProjectsMsg) {
            if (visibleCount === 0) {
                noProjectsMsg.classList.remove('d-none');
            } else {
                noProjectsMsg.classList.add('d-none');
            }
        }
    }

    window.resetProjectFilters = function() {
        if (searchInput) searchInput.value = '';
        searchQuery = '';
        activeFilter = 'all';

        filterButtons.forEach(b => {
            if (b.getAttribute('data-filter') === 'all') {
                b.classList.remove('btn-outline-secondary');
                b.classList.add('btn-primary', 'active');
            } else {
                b.classList.remove('btn-primary', 'active');
                b.classList.add('btn-outline-secondary');
            }
        });

        applyFilters();
    };
}
