package main

import (
	"context"
	"fmt"
	"log"

	empty "github.com/golang/protobuf/ptypes/empty"
	discord "github.com/roleypoly/rpc/discord"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	transport := grpc.WithTransportCredentials(credentials.NewTLS(nil))
	grpcClient, err := grpc.Dial("localhost:6777", transport)
	if err != nil {
		log.Fatalln(err)
	}

	dd := discord.NewDiscordClient(grpcClient)

	guildList, err := dd.ListGuilds(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, guild := range guildList.Guilds {
		fmt.Printf("%s - %d members\n", guild.Name, guild.MemberCount)
	}
}
