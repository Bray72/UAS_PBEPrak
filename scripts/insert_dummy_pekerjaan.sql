-- Updated to use alumni NIM references instead of hardcoded IDs
-- Insert dummy data untuk pekerjaan_alumni (100+ records)
-- Menggunakan subquery untuk mendapatkan alumni_id berdasarkan NIM yang baru dibuat

INSERT INTO pekerjaan_alumni (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, created_at, updated_at) VALUES

-- Alumni dengan NIM 434231101 - 2 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231101'), 'PT Tokopedia', 'Software Engineer', 'E-commerce', 'Jakarta', '15000000-20000000', '2023-07-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231101'), 'PT Gojek Indonesia', 'Junior Developer', 'Technology', 'Jakarta', '12000000-15000000', '2023-01-01', '2023-06-30', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231102 - 2 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231102'), 'PT Bank Central Asia', 'System Analyst', 'Banking', 'Jakarta', '18000000-25000000', '2023-08-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231102'), 'PT Telkom Indonesia', 'IT Support', 'Telecommunications', 'Jakarta', '10000000-12000000', '2023-02-01', '2023-07-31', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231103 - 2 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231103'), 'PT Shopee Indonesia', 'Backend Developer', 'E-commerce', 'Bandung', '16000000-22000000', '2023-09-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231103'), 'PT Traveloka', 'Software Developer', 'Travel Technology', 'Bandung', '14000000-18000000', '2023-03-01', '2023-08-31', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231104 - 1 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231104'), 'PT Bukalapak', 'Frontend Developer', 'E-commerce', 'Surabaya', '13000000-17000000', '2024-01-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231105 - 2 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231105'), 'PT Grab Indonesia', 'Data Analyst', 'Transportation Technology', 'Yogyakarta', '14000000-19000000', '2024-02-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231105'), 'PT Blibli.com', 'Junior Analyst', 'E-commerce', 'Yogyakarta', '11000000-14000000', '2023-08-01', '2024-01-31', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231106 - 1 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231106'), 'PT Astra International', 'IT Specialist', 'Automotive', 'Medan', '15000000-20000000', '2024-03-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231107 - 3 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231107'), 'PT Google Indonesia', 'Senior Software Engineer', 'Technology', 'Jakarta', '35000000-45000000', '2023-01-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231107'), 'PT Microsoft Indonesia', 'Software Engineer', 'Technology', 'Jakarta', '25000000-30000000', '2022-07-01', '2022-12-31', 'selesai', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231107'), 'PT Oracle Indonesia', 'Junior Developer', 'Technology', 'Jakarta', '18000000-22000000', '2022-01-01', '2022-06-30', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231108 - 2 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231108'), 'PT Mandiri Sekuritas', 'Business Analyst', 'Financial Services', 'Surabaya', '16000000-21000000', '2022-08-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231108'), 'PT Pegadaian', 'IT Analyst', 'Financial Services', 'Surabaya', '13000000-16000000', '2022-02-01', '2022-07-31', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231109 - 2 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231109'), 'PT Pertamina', 'System Administrator', 'Oil & Gas', 'Bandung', '17000000-23000000', '2022-09-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231109'), 'PT PLN', 'Network Engineer', 'Energy', 'Bandung', '14000000-18000000', '2022-03-01', '2022-08-31', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231110 - 1 pekerjaan
((SELECT id FROM alumni WHERE nim = '434231110'), 'PT Startup Semarang', 'Intern Developer', 'Technology', 'Semarang', '3000000-5000000', '2024-06-01', '2024-08-31', 'selesai', NOW(), NOW()),

-- Alumni dengan NIM 434231111-120
((SELECT id FROM alumni WHERE nim = '434231111'), 'PT Unilever Indonesia', 'IT Developer', 'Consumer Goods', 'Surabaya', '16000000-21000000', '2024-01-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231112'), 'PT Indofood', 'System Analyst', 'Food & Beverage', 'Makassar', '14000000-18000000', '2024-02-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231113'), 'PT Samsung Electronics', 'Software Engineer', 'Electronics', 'Jakarta', '20000000-28000000', '2021-08-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231114'), 'PT Xiaomi Indonesia', 'Mobile Developer', 'Electronics', 'Surabaya', '18000000-24000000', '2021-09-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231115'), 'PT Garuda Indonesia', 'IT Support', 'Aviation', 'Palembang', '12000000-16000000', '2021-10-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231116-125
((SELECT id FROM alumni WHERE nim = '434231116'), 'PT Agoda', 'Frontend Developer', 'Travel Technology', 'Denpasar', '15000000-20000000', '2025-01-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231117'), 'PT Tiket.com', 'Backend Developer', 'Travel Technology', 'Surabaya', '14000000-19000000', '2025-02-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231118'), 'PT Ruangguru', 'Full Stack Developer', 'Education Technology', 'Balikpapan', '16000000-22000000', '2025-03-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231119'), 'PT Zenius Education', 'Software Engineer', 'Education Technology', 'Jakarta', '17000000-23000000', '2020-07-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231120'), 'PT Halodoc', 'Mobile Developer', 'Healthcare Technology', 'Surabaya', '18000000-25000000', '2020-08-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231121-130
((SELECT id FROM alumni WHERE nim = '434231121'), 'PT Alodokter', 'Data Scientist', 'Healthcare Technology', 'Padang', '19000000-26000000', '2020-09-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231122'), 'PT KitaBisa', 'DevOps Engineer', 'Social Technology', 'Manado', '16000000-21000000', '2026-01-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231123'), 'PT Flip', 'Backend Engineer', 'Fintech', 'Surabaya', '17000000-24000000', '2026-02-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231124'), 'PT Dana Indonesia', 'Mobile Engineer', 'Fintech', 'Pontianak', '18000000-25000000', '2026-03-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231125'), 'PT OVO', 'Senior Developer', 'Fintech', 'Jakarta', '22000000-30000000', '2019-06-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231126-135
((SELECT id FROM alumni WHERE nim = '434231126'), 'PT LinkAja', 'Software Architect', 'Fintech', 'Surabaya', '25000000-35000000', '2019-07-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231127'), 'PT Jenius', 'Lead Developer', 'Banking Technology', 'Makassar', '24000000-32000000', '2019-08-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231128'), 'PT Kredivo', 'Machine Learning Engineer', 'Fintech', 'Batam', '20000000-28000000', '2027-01-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231129'), 'PT Akulaku', 'Data Engineer', 'Fintech', 'Surabaya', '19000000-26000000', '2027-02-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231130'), 'PT Investree', 'Blockchain Developer', 'Fintech', 'Solo', '21000000-29000000', '2027-03-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231131-140
((SELECT id FROM alumni WHERE nim = '434231131'), 'PT Amartha', 'Full Stack Engineer', 'Fintech', 'Jakarta', '18000000-25000000', '2018-05-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231132'), 'PT Modalku', 'DevOps Specialist', 'Fintech', 'Surabaya', '17000000-23000000', '2018-06-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231133'), 'PT Koinworks', 'Security Engineer', 'Fintech', 'Yogyakarta', '19000000-27000000', '2018-07-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231134'), 'PT Payfazz', 'Mobile App Developer', 'Fintech', 'Bandung', '16000000-22000000', '2017-04-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231135'), 'PT Doku', 'Payment System Developer', 'Fintech', 'Surabaya', '18000000-24000000', '2017-05-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231136-145
((SELECT id FROM alumni WHERE nim = '434231136'), 'PT Midtrans', 'API Developer', 'Payment Technology', 'Medan', '17000000-23000000', '2017-06-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231137'), 'PT Xendit', 'Senior Backend Engineer', 'Payment Technology', 'Jakarta', '23000000-31000000', '2016-03-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231138'), 'PT Faspay', 'Integration Specialist', 'Payment Technology', 'Surabaya', '15000000-20000000', '2016-04-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231139'), 'PT Nicepay', 'Technical Lead', 'Payment Technology', 'Semarang', '21000000-28000000', '2016-05-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231140'), 'PT Espay', 'Software Developer', 'Payment Technology', 'Palembang', '14000000-19000000', '2015-02-01', NULL, 'aktif', NOW(), NOW()),

-- Alumni dengan NIM 434231141-150
((SELECT id FROM alumni WHERE nim = '434231141'), 'PT Veritrans', 'QA Engineer', 'Payment Technology', 'Surabaya', '13000000-17000000', '2015-03-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231142'), 'PT Kartuku', 'Mobile Payment Developer', 'Payment Technology', 'Pekanbaru', '16000000-21000000', '2015-04-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231143'), 'PT Cashlez', 'Fintech Developer', 'Payment Technology', 'Jakarta', '17000000-23000000', '2014-01-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231144'), 'PT Artajasa', 'System Integration', 'Payment Technology', 'Surabaya', '15000000-20000000', '2014-02-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231145'), 'PT Finnet Indonesia', 'Network Developer', 'Financial Technology', 'Banjarmasin', '14000000-18000000', '2014-03-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231146'), 'PT Rintis', 'Startup Developer', 'Technology', 'Denpasar', '12000000-16000000', '2013-12-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231147'), 'PT Tech in Asia', 'Content Management System Developer', 'Media Technology', 'Surabaya', '13000000-17000000', '2013-11-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231148'), 'PT Daily Social', 'Web Developer', 'Digital Media', 'Samarinda', '11000000-15000000', '2013-10-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231149'), 'PT Techinasia', 'Junior Full Stack', 'Technology Media', 'Jakarta', '14000000-18000000', '2012-09-01', NULL, 'aktif', NOW(), NOW()),
((SELECT id FROM alumni WHERE nim = '434231150'), 'PT Startup Indonesia', 'Software Engineer', 'Technology', 'Surabaya', '15000000-20000000', '2012-08-01', NULL, 'aktif', NOW(), NOW());
