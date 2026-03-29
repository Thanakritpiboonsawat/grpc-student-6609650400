package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//เรียก GetStudent ก่อน
	res, err := client.GetStudent(ctx, &pb.StudentRequest{
		Id: 101,
	})
	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Println("Single Student:")
	log.Printf("ID: %d | Name: %s | Major: %s | Email: %s | Phone: %s",
		res.Id,
		res.Name,
		res.Major,
		res.Email,
		res.Phone,
	)

	//แล้วค่อยเรียก ListStudents
	listRes, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	log.Println("Student List:")
	for _, student := range listRes.Students {
		log.Printf("ID: %d | Name: %s | Major: %s | Email: %s | Phone: %s",
			student.Id,
			student.Name,
			student.Major,
			student.Email,
			student.Phone,
		)
	}
}
