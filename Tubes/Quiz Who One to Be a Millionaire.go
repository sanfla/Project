package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NMAX int = 1000

type question struct {
	pertanyaan, kunci          string
	opsiA, opsiB, opsiC, opsiD string
	salah, benar               int
}
type questions [NMAX]question

type user struct {
	name, username, password string
	skor                     float64
}
type users [NMAX]user

func main() {
	var User users
	var Q questions
	var password string
	var nQuest int = 0
	var nUser int = 0
	header_demo()
	passwordAdmin(&password)
	header_mulai()
	menu_utama(&Q, &nQuest, &User, &nUser, &password)
}

func header_demo() {
	/* Menampilkan header pembatas. Tampilan berupa judul tugas besar dan identitas pembuat. */
	fmt.Println()
	fmt.Println(" * ------------------------------------------- * ")
	fmt.Println(" *          Aplikasi Demo Tugas Besar          * ")
	fmt.Println(" *     Created by Bagas Eko Tjahyono Putro     * ")
	fmt.Println(" *           & Salma Rizki Nurfauziah          * ")
	fmt.Println(" *          Algoritma Pemrograman 2023         * ")
	fmt.Println(" * ------------------------------------------- * ")
}

func passwordAdmin(password *string) {
	/* I.S. -
	   proses : menerima masukan dari pengguna berupa password admin. Password ini akan digunakan untuk memasuki menu admin.
	   F.S. Menampilkan header kunci admin dan meminta masukan. Inputan pengguna akan tersimpan sebagai password yang kemudian value-nya
	        dilempar kembali ke func main.
	*/
	fmt.Println("\n-------------------------------------------")
	fmt.Println("             [ KUNCI ADMIN ]")
	fmt.Println("-------------------------------------------")
	fmt.Print("Masukkan Password Admin : ")
	fmt.Scan(&*password)
}

func header_mulai() {
	/* Menampilkan header memulai aplikasi. Tampilan berupa judul dan quote singkat aplikasi. */
	var y rune
	fmt.Println()
	fmt.Println("============ WELCOME ============")
	fmt.Println("   WHO ONE TO BE A MILLIONAIRE   ")
	fmt.Println("=================================")
	fmt.Println("    Do you have what it takes    ")
	fmt.Println("      to be the champion and     ")
	fmt.Println("        WON ONE MILLIONS?        ")
	fmt.Println("=================================")
	fmt.Println()
	fmt.Print("Press y to continue....")
	fmt.Scan(&y)
}

func header_menu() {
	/* Menampilkan header menu utama. Tampilan berupa pilihan angka serta deskripsi singkat interaksi yang dapat dilakukannya. */
	header_demo()
	fmt.Println("=====================")
	fmt.Println("         MENU        ")
	fmt.Println("=====================")
	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Lihat Skor")
	fmt.Println("4. Admin")
	fmt.Println("5. Exit")
	fmt.Println("=====================")
}

func menu_utama(Quest *questions, nQuest *int, User *users, nUser *int, password *string) {
	/* I.S. -
	   proses: menerima masukan dari keyboard berupa pilihan pengguna. Apabila masukan tidak memenuhi syarat (memilih angka di luar 1-5), maka pengguna diminta mengulang masukan.
	   F.S. Memanggil function/procedure spesifik sesuai korespondensi masukan. Apabila masukan berupa angka 5, maka akan menghentikan aplikasi.
	*/
	var banyakSoal int
	var choice string
	header_menu()
	fmt.Print("Pilihan Anda (1/2/3/4/5): ")
	fmt.Scan(&choice)
	for choice != "5" {
		for choice != "1" && choice != "2" && choice != "3" && choice != "4" && choice != "5" {
			fmt.Print("Pilihan Anda (1/2/3/4/5): ")
			fmt.Scan(&choice)
		}
		if choice == "1" {
			registrasi(&*User, &*nUser)
		} else if choice == "2" {
			login(&*Quest, &*nQuest, &*User, &*nUser, banyakSoal)
		} else if choice == "3" {
			lihat_skor(*User, *nUser)
		} else if choice == "4" {
			admin(&*Quest, &*nQuest, &*User, &*nUser, &*password, &banyakSoal)
		}
		header_menu()
		fmt.Print("Pilihan Anda (1/2/3/4/5): ")
		fmt.Scan(&choice)
	}
	fmt.Println("\n~ Thank you for playing! ^^ ~")
}

func cek_username(u users, nUsers int, check string) int {
	/* I.S. Terdefinisi array u sejumlah nUsers data serta sebuah string untuk dibandingkan.
	   F.S. Mengembalikan indeks apabila username yang sudah tersimpan dengan string yang disandingkan sama, -1 apabila tidak ada kesamaan.
	*/
	for i := 0; i < nUsers; i++ {
		if check == u[i].username {
			return i
		}
	}
	return -1
}

func registrasi(u *users, nUser *int) {
	/* I.S. -
	   proses: menerima masukan dari keyboard berupa data nama,username,dan password peserta. Data dapat disimpan di penyimpanan array u.
	   F.S. array u berisi sejumlah nUser data. nUser mungkin tidak bertambah.
	*/
	var u_temp user
	var choice string
	header_demo()
	fmt.Println("=====================")
	fmt.Println("     REGISTRASI      ")
	fmt.Println("=====================")
	fmt.Print("Name : ")
	fmt.Scan(&u_temp.name)
	fmt.Print("Username : ")
	fmt.Scan(&u_temp.username)
	fmt.Print("Password : ")
	fmt.Scan(&u_temp.password)
	fmt.Println("------------------------")
	fmt.Println("1.Simpan  2.Batal")
	fmt.Print("Pilihan Anda (1/2): ")
	fmt.Scan(&choice)
	for choice != "1" && choice != "2" {
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&choice)
	}

	if choice == "1" {
		if cek_username(*u, *nUser, u_temp.username) != -1 {
			fmt.Println("\n[ Maaf, username sudah digunakan :( ]")
		} else {
			u[*nUser] = u_temp
			*nUser++
			fmt.Println("\n[ Registrasi Berhasil ^^ ]")
		}
	} else {
		fmt.Println("\n[ Membatalkan Registrasi.... ]")
	}
}

func ascendingSortSkor(v *users, n int) {
	var pass, idx, i int
	var temp user
	pass = 1
	for pass <= n {
		idx = pass - 1
		i = pass
		for i < n {
			if v[idx].skor > v[i].skor {
				idx = i
			}
			i++
		}
		temp = v[pass-1]
		v[pass-1] = v[idx]
		v[idx] = temp
		pass++
	}
}

func descendingSortSkor(v *users, n int) {
	var pass, idx, i int
	var temp user
	pass = 1
	for pass <= n {
		idx = pass - 1
		i = pass
		for i < n {
			if v[idx].skor < v[i].skor {
				idx = i
			}
			i++
		}
		temp = v[pass-1]
		v[pass-1] = v[idx]
		v[idx] = temp
		pass++
	}
}

func lihat_skor(u users, n int) {
	/* I.S. Terdefinisi array u sejumlah n data.
	   F.S. Menampilkan nama, username, serta skor peserta secara vertikal dengan data-data yang ditampilkan menyamping.
	*/
	var choice string
	var nama string
	for choice != "4" {
		header_demo()
		fmt.Println("=====================")
		fmt.Println("     SKOR PEMAIN     ")
		fmt.Println("=====================")
		for i := 0; i < n; i++ {
			fmt.Print(i+1, ". ", u[i].name, " (@", u[i].username, ") [ Skor : ", u[i].skor, " ]\n")
		}
		fmt.Println("------------------")
		fmt.Println("1. Ascending by Score  2. Descending by Score  3. Cari Pemain   4. Keluar")
		fmt.Print("Pilihan Anda (1/2/3/4): ")
		fmt.Scan(&choice)
		for choice != "1" && choice != "2" && choice != "3" && choice != "4" {
			fmt.Print("Pilihan Anda (1/2/3/4): ")
			fmt.Scan(&choice)
		}
		if choice == "1" {
			ascendingSortSkor(&u, n)
		} else if choice == "2" {
			descendingSortSkor(&u, n)
		} else if choice == "3" {
			fmt.Print("Masukkan username pemain: ")
			fmt.Scan(&nama)
			sortingPeserta(&u, n)
			idx := binarySearchPeserta(u, n, nama)
			if idx == -1 {
				fmt.Println("[ Maaf, tidak ada pemain dengan username tersebut :( ]")
			} else {
				fmt.Println("==== PESERTA ====")
				fmt.Print(u[idx].name, " (@", u[idx].username, ") [ Skor : ", u[idx].skor, " ]\n")
			}
		}
	}
}

func sortingPeserta(U *users, n int) {
	var pass, i int
	var temp user
	pass = 1
	for pass < n {
		temp = U[pass]
		i = pass
		for i > 0 && temp.username < U[i-1].username {
			U[i] = U[i-1]
			i--
		}
		U[i] = temp
		pass++
	}
}

func binarySearchPeserta(U users, n int, x string) int {
	var mid, left, right, idx int
	left = 0
	right = n
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if x < U[mid].username {
			right = mid - 1
		} else if x > U[mid].username {
			left = mid + 1
		} else {
			idx = mid
			return idx
		}
	}
	return idx
}

func admin(Q *questions, nQuest *int, U *users, nUser *int, pass *string, banyakSoal *int) {
	var choice string
	var isi string
	fmt.Println("\n========= LOGIN ==========")
	fmt.Print("Password : ")
	fmt.Scan(&isi)
	if isi == *pass {
		for choice != "5" {
			header_demo()
			fmt.Println("=====================")
			fmt.Println("     MENU ADMIN      ")
			fmt.Println("=====================")
			fmt.Println("1. Tambah soal")
			fmt.Println("2. Tinjau soal")
			fmt.Println("3. Lihat soal tersusah/termudah")
			fmt.Println("4. Setting Banyak Soal Quiz")
			fmt.Println("5. Kembali")
			fmt.Println("------------------")
			fmt.Print("Pilihan Anda (1/2/3/4/5): ")
			fmt.Scan(&choice)
			for choice != "1" && choice != "2" && choice != "3" && choice != "4" && choice != "5" {
				fmt.Print("Pilihan Anda (1/2/3/4/5): ")
				fmt.Scan(&choice)
			}
			if choice == "1" {
				tambahSoal(&*Q, &*nQuest, &*banyakSoal)
			} else if choice == "2" {
				tinjauSoal(&*Q, &*nQuest)
			} else if choice == "3" {
				tinjauSoalTermudahTersulit(*Q, *nQuest)
			} else if choice == "4" {
				gantiBanyakSoal(&*banyakSoal, *nQuest)
			}
		}
	} else {
		fmt.Println("\n[ Sorry, password incorrect :( ]")
	}
}

func gantiBanyakSoal(banyakSoal *int, nQuest int) {
	var choice string
	var temp int
	fmt.Println("Berapa banyak soal peserta akan kerjakan (Default sebanyak jumlah bank soal):", *banyakSoal)
	fmt.Println("Ganti?")
	fmt.Println("------------------")
	fmt.Println("1. Ya    2. Tidak")
	fmt.Print("Pilihan Anda (1/2): ")
	fmt.Scan(&choice)
	for choice != "1" && choice != "2" {
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&choice)
	}
	if choice == "1" {
		fmt.Print("\nSet berapa banyak soal peserta akan kerjakan (Tak boleh lebih dari banyak bank soal): ")
		fmt.Scan(&temp)
		for temp > nQuest {
			fmt.Print("Set berapa banyak soal peserta akan kerjakan (Tak boleh lebih dari banyak bank soal): ")
			fmt.Scan(&temp)
		}
		fmt.Println("------------------")
		fmt.Println("1. Simpan    2. Batal")
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&choice)
		for choice != "1" && choice != "2" {
			fmt.Print("Pilihan Anda (1/2): ")
			fmt.Scan(&choice)
		}
		if choice == "1" {
			*banyakSoal = temp
			fmt.Println("\n[ Banyak soal peserta berhasil diganti ^^ ]")
		}
	}
}

func inputSoal(q *question) {
	fmt.Println(" Pertanyaan: (gunakan underscore untuk pengganti spasi)")
	fmt.Scan(&q.pertanyaan)
	fmt.Print("Opsi A: ")
	fmt.Scan(&q.opsiA)
	fmt.Print("Opsi B: ")
	fmt.Scan(&q.opsiB)
	fmt.Print("Opsi C: ")
	fmt.Scan(&q.opsiC)
	fmt.Print("Opsi D: ")
	fmt.Scan(&q.opsiD)
	fmt.Print("Kunci Jawaban (A/B/C/D): ")
	fmt.Scan(&q.kunci)
	for q.kunci != "A" && q.kunci != "B" && q.kunci != "C" && q.kunci != "D" {
		fmt.Print("Kunci Jawaban (A/B/C/D): ")
		fmt.Scan(&q.kunci)
	}
}

func tambahSoal(q *questions, n *int, banyakSoal *int) {
	var choice string
	var q_temp question
	header_demo()
	fmt.Println("=====================")
	fmt.Println("     TAMBAH SOAL     ")
	fmt.Println("=====================")
	inputSoal(&q_temp)
	fmt.Println("------------------------------------------------------")
	fmt.Println("1.Simpan  2.Batal")
	fmt.Print("Pilihan Anda (1/2): ")
	fmt.Scan(&choice)
	for choice != "1" && choice != "2" {
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&choice)
	}
	if choice == "1" {
		q[*n] = q_temp
		*n++
		*banyakSoal = *n
		fmt.Println("[ Pertanyaan berhasil disimpan ^^ ]")
		fmt.Println("[ Settingan banyak soal peserta akan terupdate menjadi default ]")
	} else {
		fmt.Println("[ Pertanyaan tidak berhasil disimpan.... ]")
	}
}

func tinjauSoal(Q *questions, nQuest *int) {
	var i int = 0
	var choice string
	header_demo()
	fmt.Println("======================")
	fmt.Println("   MENU TINJAU SOAL   ")
	fmt.Println("======================")
	for i < *nQuest && choice != "5" {
		fmt.Println("-----------------------------")
		tampilSoal(*Q, i)
		fmt.Println("Berapa kali terjawab benar:", Q[i].benar)
		fmt.Println("Berapa kali terjawab salah:", Q[i].salah)
		fmt.Println("-----------------------------")
		fmt.Println("1. Next,  2. Previous,  3. Ubah,  4. Hapus,  5. Kembali")
		fmt.Print("Pilihan Anda (1/2/3/4/5): ")
		fmt.Scan(&choice)
		for choice != "1" && choice != "2" && choice != "3" && choice != "4" && choice != "5" {
			fmt.Print("Pilihan Anda (1/2/3/4/5): ")
			fmt.Scan(&choice)
		}
		if choice == "1" {
			i++
		} else if choice == "2" {
			if i > 0 {
				i--
			}
		} else if choice == "3" {
			fmt.Println("Perubahan:")
			ubahSoal(&*Q, i)
			i++
			fmt.Println("-----------------------------")
		} else if choice == "4" {
			hapusSoal(&*Q, &*nQuest, i)
		}
		fmt.Println()
	}
}

func ubahSoal(q *questions, i int) {
	var choice string
	var q_temp question
	inputSoal(&q_temp)
	fmt.Println("------------------------------------------------------")
	fmt.Println("1.Simpan  2.Batal")
	fmt.Print("Pilihan Anda (1/2): ")
	fmt.Scan(&choice)
	for choice != "1" && choice != "2" {
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&choice)
	}
	if choice == "1" {
		q[i] = q_temp
		fmt.Println("[ Pertanyaan berhasil disimpan ^^ ]")
	}
}

func hapusSoal(q *questions, n *int, j int) {
	var choice string
	fmt.Println("1.Hapus  2.Batal")
	fmt.Print("Pilihan Anda (1/2): ")
	fmt.Scan(&choice)
	for choice != "1" && choice != "2" {
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&choice)
	}
	if choice == "1" {
		for j <= *n-2 {
			q[j] = q[j+1]
			j++
		}
		*n--
		fmt.Println("[ Pertanyaan Berhasil Dihapus ^^ ]")
	}
}

func soalTermudah(q *questions, n int) {
	var pass, i int
	var temp question
	pass = 1
	for pass < n {
		temp = q[pass]
		i = pass
		for i > 0 && temp.benar > q[i-1].benar {
			q[i] = q[i-1]
			i--
		}
		q[i] = temp
		pass++
	}
}

func soalTersulit(q *questions, n int) {
	var pass, i int
	var temp question
	pass = 1
	for pass < n {
		temp = q[pass]
		i = pass
		for i > 0 && temp.salah > q[i-1].salah {
			q[i] = q[i-1]
			i--
		}
		q[i] = temp
		pass++
	}
}

func tampilSoal(Q questions, i int) {
	fmt.Println("Soal: ", i+1)
	fmt.Println("Pertanyaan: ", Q[i].pertanyaan)
	fmt.Println("A.", Q[i].opsiA, "        ", "B.", Q[i].opsiB)
	fmt.Println("C.", Q[i].opsiC, "        ", "D.", Q[i].opsiD)
	fmt.Println("Kunci jawaban: ", Q[i].kunci)
}

func tinjauSoalTermudahTersulit(Q questions, nQuest int) {
	var i, j int
	var choice1, choice2 string
	for choice1 != "3" {
		header_demo()
		fmt.Println("======================")
		fmt.Println("      TINJAU SOAL     ")
		fmt.Println("======================")
		fmt.Println("1. Lihat soal termudah")
		fmt.Println("2. Lihat soal tersulit")
		fmt.Println("3. Kembali")
		fmt.Println("-----------------------------")
		fmt.Print("Pilihan Anda (1/2/3): ")
		fmt.Scan(&choice1)
		for choice1 != "1" && choice1 != "2" && choice1 != "3" {
			fmt.Print("Pilihan Anda (1/2/3): ")
			fmt.Scan(&choice1)
		}
		if choice1 == "1" {
			i = 0
			j = 0
			soalTermudah(&Q, nQuest)
			if Q[i].benar > 0 {
				fmt.Println("\n====== TERMUDAH ======")
			}
			for j < 5 && Q[i].benar > 0 && choice2 != "3" {
				if Q[i].benar > Q[i].salah {
					tampilSoal(Q, j)
					fmt.Println("Berapa kali terjawab benar:", Q[i].benar)
					fmt.Println("Berapa kali terjawab salah:", Q[i].salah)
					fmt.Println("-----------------------------")
					fmt.Println("1. Next,  2. Previous,  3. Kembali")
					fmt.Print("Pilihan Anda (1/2/3): ")
					fmt.Scan(&choice2)
					for choice2 != "1" && choice2 != "2" && choice2 != "3" {
						fmt.Print("Pilihan Anda (1/2/3): ")
						fmt.Scan(&choice2)
					}
					fmt.Println("-----------------------------")
					if choice2 == "1" {
						j++
						i++
					} else if choice2 == "2" {
						if i > 0 {
							i--
							j--
						}
					}
				} else {
					i++
				}
			}
		} else if choice1 == "2" {
			i = 0
			j = 0
			soalTersulit(&Q, nQuest)
			if Q[i].salah > 0 {
				fmt.Println("\n====== TERSULIT ======")
			}
			for j < 5 && Q[i].salah > 0 && choice2 != "3" {
				if Q[i].salah > Q[i].benar {
					tampilSoal(Q, j)
					fmt.Println("Berapa kali terjawab salah:", Q[i].salah)
					fmt.Println("Berapa kali terjawab benar:", Q[i].benar)
					fmt.Println("-----------------------------")
					fmt.Println("1. Next,  2. Previous,  3. Kembali")
					fmt.Print("Pilihan Anda (1/2/3): ")
					fmt.Scan(&choice2)
					for choice2 != "1" && choice2 != "2" && choice2 != "3" {
						fmt.Print("Pilihan Anda (1/2/3): ")
						fmt.Scan(&choice2)
					}
					fmt.Println("-----------------------------")
					if choice2 == "1" {
						i++
						j++
					} else if choice2 == "2" {
						if i > 0 {
							i--
							j--
						}
					}
				} else {
					i++
				}
			}
		}
	}
}

func login(Q *questions, nQuest *int, U *users, nUser *int, banyakSoal int) {
	var uss, pass, choice string
	var y rune
	header_demo()
	fmt.Println("========= LOGIN ==========")
	fmt.Print("Username: ")
	fmt.Scan(&uss)
	fmt.Print("Password: ")
	fmt.Scan(&pass)
	fmt.Println("-----------------------------")
	fmt.Println("1.Login  2.Batal")
	fmt.Print("Pilihan Anda (1/2): ")
	fmt.Scan(&choice)
	for choice != "1" && choice != "2" {
		fmt.Print("Pilihan Anda (1/2): ")
		fmt.Scan(&choice)
	}
	if choice == "1" {
		check := cek_username(*U, *nUser, uss)
		if check != -1 && U[check].password == pass {
			fmt.Println("[ Login Berhasil.... ]")
			header_demo()
			fmt.Println("====================")
			fmt.Println("    MENU PESERTA    ")
			fmt.Println("====================")
			fmt.Print("Selamat datang, ", U[check].name, ".\n\n")
			fmt.Println("===== MEMULAI QUIZ =====")
			fmt.Print("Press y to continue....")
			fmt.Scan(&y)

			soal(&*Q, *nQuest, &*U, check, banyakSoal)
		} else {
			fmt.Println("[ Username atau password yang Anda masukan salah :( ]")
		}
	}
}

func soal(Q *questions, nQuest int, U *users, nomor, banyakSoal int) {
	var i, j, temp, score int
	var kunci string
	for i = 0; i < banyakSoal; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		j = r1.Intn(nQuest)
		for j == temp {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			j = r1.Intn(nQuest)
		}
		fmt.Println("\nPertanyaan", i+1, ":", Q[j].pertanyaan)
		fmt.Println("A.", Q[j].opsiA, "   ", "B.", Q[j].opsiB)
		fmt.Println("C.", Q[j].opsiC, "   ", "D.", Q[j].opsiD)
		fmt.Println("-----------------------------------------------")
		fmt.Print("Jawaban Anda (A/B/C/D): ")
		fmt.Scan(&kunci)
		for kunci != "A" && kunci != "B" && kunci != "C" && kunci != "D" {
			fmt.Print("Jawaban Anda (A/B/C/D): ")
			fmt.Scan(&kunci)
		}
		fmt.Println()
		if kunci == Q[j].kunci {
			score++
			Q[j].benar++
			fmt.Println("DING DING, JAWABAN BENAR ^_^ ")
			fmt.Println("Jawabannya adalah", Q[j].kunci)
			fmt.Println()
		} else {
			Q[j].salah++
			fmt.Println("TET TOT, JAWABAN SALAH!")
			fmt.Println("Jawabannya adalah", Q[j].kunci)
			fmt.Println()
		}
		temp = j
	}
	fmt.Println("-----------------------------------------------")
	fmt.Println("Menghitung....\n")
	U[nomor].skor = float64(score) * 100 / float64(i)
	fmt.Println("[ SKOR AKHIR ANDA :", U[nomor].skor, "]\n")
}


