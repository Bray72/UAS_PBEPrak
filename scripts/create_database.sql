-- Create database and tables for alumni management system
CREATE DATABASE alumni_db;

-- Connect to the database
\c alumni_db;

-- Adding users table for authentication
-- Create Users table for authentication
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(20) DEFAULT 'user' CHECK (role IN ('admin', 'user')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample users (password: "123456")
INSERT INTO users (username, email, password_hash, role) VALUES
('admin', 'admin@university.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'admin'),
('user1', 'user1@university.com', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'user');

-- Create Alumni table
CREATE TABLE alumni (
    id SERIAL PRIMARY KEY,
    nim VARCHAR(20) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    jurusan VARCHAR(50) NOT NULL,
    angkatan INTEGER NOT NULL,
    tahun_lulus INTEGER NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    no_telepon VARCHAR(15),
    alamat TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Pekerjaan Alumni table
CREATE TABLE pekerjaan_alumni (
    id SERIAL PRIMARY KEY,
    alumni_id INTEGER NOT NULL,
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
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (alumni_id) REFERENCES alumni(id) ON DELETE CASCADE
);

-- Insert sample data
INSERT INTO alumni (nim, nama, jurusan, angkatan, tahun_lulus, email, no_telepon, alamat) VALUES
('2021001', 'Moh Nasrul Aziz', 'Informatika', 2021, 2025, 'aziz.mnasrul@gmail.com', '081234567890', 'Surabaya'),
('2021002', 'Anugrah Anang Prastyo', 'Sistem Informasi', 2021, 2025, 'anugrah.anang@gmail.com', '081234567891', 'Malang'),
('2020001', 'Christian Dimas Renggana', 'Informatika', 2020, 2024, 'christian.dimas@gmail.com', '081234567892', 'Jakarta');

-- Insert sample job data
INSERT INTO pekerjaan_alumni (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range, tanggal_mulai_kerja, status_pekerjaan, deskripsi_pekerjaan) VALUES
(1, 'PT Tech Indonesia', 'Software Engineer', 'Technology', 'Jakarta', '8-12 juta', '2025-01-15', 'aktif', 'Mengembangkan aplikasi web menggunakan Go dan React'),
(2, 'CV Digital Solutions', 'Frontend Developer', 'Technology', 'Surabaya', '6-10 juta', '2025-02-01', 'aktif', 'Membuat antarmuka pengguna yang responsif'),
(3, 'PT Startup Nusantara', 'Full Stack Developer', 'Technology', 'Bandung', '10-15 juta', '2024-06-01', 'aktif', 'Mengembangkan sistem informasi perusahaan');

-- psql -U postgres -d alumni_db
