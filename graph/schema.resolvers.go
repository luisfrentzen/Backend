package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Backend/graph/generated"
	"Backend/graph/model"
	"context"
	"errors"
	"fmt"
	"log"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	user := model.User{
		ID:         input.ID,
		Name:       input.Name,
		Profilepic: input.Profilepic,
		Premium:    input.Premium,
	}

	log.Println("Inserting User")

	_, err := r.DB.Model(&user).Insert()


	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert new user failed")
	} else {
		log.Println("Create User Succeed")
	}

	return &user, nil
	//panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.NewUser) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("User not found!")
	}

	user.Name = input.Name

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update user failed")
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("User not found!")
	}

	_, deleteErr := r.DB.Model(&user).Where("id = ?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete user failed")
	}

	return true, nil
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	var user model.User

	fkErr := r.DB.Model(&user).Where("id = ?", input.Userid).First()

	if fkErr != nil {
		fmt.Println(fkErr)
		return nil, errors.New("User not found")
	}

	video := model.Video{
		Title:       input.Title,
		URL:         input.URL,
		Userid:      input.Userid,
		Desc:        input.Desc,
		Location:    input.Location,
		Visibility:  input.Visibility,
		Thumbnail:   input.Thumbnail,
		Playlist:    input.Playlist,
		Restriction: input.Restriction,
		Category:    input.Category,
		Like:        input.Like,
		Disilike:    input.Disilike,
		View:        input.View,
	}

	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Insert new video failed")
	} else {
		log.Println("Create Video Succeed")
	}

	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input *model.NewVideo) (*model.Video, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("Video not found!")
	}

	video.Title = input.Title
	video.URL = input.URL

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		return nil, errors.New("Update video failed")
	}

	return &video, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	var video model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("Video not found!")
	}

	_, deleteErr := r.DB.Model(&video).Where("id = ?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete video failed")
	}

	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Model(&users).Order("id").Select()

	if err != nil {
		return nil, errors.New("Failed to query users")
	} else {
		log.Println("Get Users Succeed")
	}

	return users, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Order("id").Select()

	if err != nil {
		return nil, errors.New("Failed to query videos")
	}

	return videos, nil
}

func (r *queryResolver) VideosByUser(ctx context.Context, userid string) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Where("userid = ?", userid).Select()

	if err != nil {
		return nil, errors.New("Failed to query videos")
	}

	return videos, nil
}

func (r *queryResolver) UserByID(ctx context.Context, userid string) ([]*model.User, error) {
	var user []*model.User

	err := r.DB.Model(&user).Where("id = ?", userid).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query user")
	} else {
		log.Println("Get User By Id Succeed")
	}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
