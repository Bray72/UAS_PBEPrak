-- SQL Commands untuk PostgreSQL (pgAdmin4)
-- Pastikan tabel pekerjaan_alumni sudah ada, jika belum jalankan CREATE TABLE ini:

CREATE TABLE IF NOT EXISTS pekerjaan_alumni (
    id SERIAL PRIMARY KEY, -- Changed AUTO_INCREMENT to SERIAL for PostgreSQL
    alumni_id INT NOT NULL,
    nama_perusahaan VARCHAR(100) NOT NULL,
    posisi_jabatan VARCHAR(100) NOT NULL,
    bidang_industri VARCHAR(50) NOT NULL,
    lokasi_kerja VARCHAR(100) NOT NULL,
    gaji_range VARCHAR(50),
    tanggal_mulai_kerja DATE NOT NULL,
    tanggal_selesai_kerja DATE,
    status_pekerjaan VARCHAR(20) DEFAULT 'aktif' CHECK (status_pekerjaan IN ('aktif', 'selesai', 'resigned')),
    deskripsi_pekerjaan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Removed ON UPDATE for PostgreSQL compatibility
    
    FOREIGN KEY (alumni_id) REFERENCES alumni(id) ON DELETE CASCADE
);

-- Added trigger function for updated_at auto-update in PostgreSQL
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_pekerjaan_alumni_updated_at 
    BEFORE UPDATE ON pekerjaan_alumni 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Insert data pekerjaan alumni berdasarkan alumni yang ada
-- Pastikan alumni dengan ID ini sudah ada di tabel alumni

INSERT INTO pekerjaan_alumni (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, status_pekerjaan, deskripsi_pekerjaan) VALUES
-- Pekerjaan untuk John Doe (alumni_id: 6)
(6, 'PT Teknologi Maju', 'Software Engineer', 'Technology', 'Jakarta Selatan', '10-15 juta', '2023-08-01', 'aktif', 'Mengembangkan aplikasi web dan mobile menggunakan teknologi terkini'),
(6, 'CV Digital Startup', 'Junior Developer', 'Technology', 'Jakarta Selatan', '6-8 juta', '2023-01-15', 'selesai', 'Pengalaman pertama sebagai developer fresh graduate'),

-- Pekerjaan untuk Budi Santoso (alumni_id: 5)
(5, 'PT Sistem Informasi Indonesia', 'Business Analyst', 'Consulting', 'Surabaya', '8-12 juta', '2024-03-01', 'aktif', 'Menganalisis kebutuhan bisnis dan merancang solusi sistem informasi'),

-- Pekerjaan untuk Ahmad (alumni_id: 4)
(4, 'PT Tech Solutions', 'Full Stack Developer', 'Technology', 'Surabaya', '12-18 juta', '2023-09-01', 'aktif', 'Mengembangkan aplikasi web full stack menggunakan Go dan React'),
(4, 'CV Web Development', 'Frontend Developer', 'Technology', 'Surabaya', '7-10 juta', '2023-02-01', 'selesai', 'Membuat antarmuka pengguna yang responsif dan user-friendly'),

-- Pekerjaan untuk Siti Nurhaliza (alumni_id: 2)
(2, 'PT Bank Digital', 'System Analyst', 'Banking', 'Jakarta', '15-20 juta', '2023-07-01', 'aktif', 'Menganalisis dan merancang sistem perbankan digital'),

-- Pekerjaan untuk Budi Santoso yang kedua (alumni_id: 3)
(3, 'PT Software House', 'Backend Developer', 'Technology', 'Bandung', '9-14 juta', '2024-01-15', 'aktif', 'Mengembangkan API dan sistem backend menggunakan Go dan PostgreSQL'),
(3, 'CV Freelance Tech', 'Freelance Developer', 'Technology', 'Bandung', '5-8 juta', '2024-06-01', 'resigned', 'Mengerjakan proyek freelance pengembangan website');

-- Query untuk melihat data yang sudah diinsert
SELECT 
    pa.id,
    a.nama as nama_alumni,
    pa.nama_perusahaan,
    pa.posisi_jabatan,
    pa.bidang_industri,
    pa.lokasi_kerja,
    pa.gaji_range,
    pa.tanggal_mulai_kerja,
    pa.status_pekerjaan,
    pa.created_at
FROM pekerjaan_alumni pa
JOIN alumni a ON pa.alumni_id = a.id
ORDER BY pa.created_at DESC;
