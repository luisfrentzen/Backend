package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Backend/graph/generated"
	"Backend/graph/model"
	"context"
	"errors"
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
		log.Println(fkErr)
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
		Restriction: input.Restriction,
		Category:    input.Category,
		Like:        input.Like,
		Disilike:    input.Disilike,
		View:        input.View,
		Channelpic:  input.Channelpic,
		Channelname: input.Channelname,
		Day:         input.Day,
		Month:       input.Month,
		Year:        input.Year,
	}

	_, err := r.DB.Model(&video).Insert()

	if err != nil {
		log.Println(err)
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
	video.Desc = input.Desc
	video.View = input.View
	video.Like = input.Like
	video.Disilike = input.Disilike

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

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	playlist := model.Playlist{
		Title:      input.Title,
		Userid:     input.Userid,
		Desc:       input.Desc,
		Visibility: input.Visibility,
		View:       input.View,
		Day:        input.Day,
		Month:      input.Month,
		Year:       input.Year,
		Videos:     input.Videos,
	}

	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert new playlist failed")
	} else {
		log.Println("Create Playlist Succeed")
	}

	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, id string, input *model.NewPlaylist) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		return nil, errors.New("Playlist not found!")
	}

	playlist.Title = input.Title
	playlist.Desc = input.Desc
	playlist.Videos = input.Videos
	playlist.Visibility = input.Visibility
	playlist.Day = input.Day
	playlist.Month = input.Month
	playlist.Year = input.Year

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) AddToPlaylist(ctx context.Context, id string, input *model.AddToPlaylist) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Playlist not found!")
	}

	if playlist.Videos == "" {
		playlist.Videos = input.Videos
	} else {
		playlist.Videos += "," + input.Videos
	}

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	return &playlist, nil
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

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Order("id").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query videos")
	}

	return playlists, nil
}

func (r *queryResolver) PlaylistsByUser(ctx context.Context, userid string) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Where("userid = ?", userid).Select()

	if err != nil {
		return nil, errors.New("Failed to query playlists")
	}

	return playlists, nil
}

func (r *queryResolver) VideosByUser(ctx context.Context, userid string) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Where("userid = ?", userid).Select()

	if err != nil {
		return nil, errors.New("Failed to query videos")
	}

	return videos, nil
}

func (r *queryResolver) VideosByCategory(ctx context.Context, category string) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Where("category = ?", category).Select()

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

func (r *queryResolver) VideoByID(ctx context.Context, id int) ([]*model.Video, error) {
	var video []*model.Video

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query video")
	} else {
		log.Println("Get Video By Id Succeed")
	}
	return video, nil
}

func (r *queryResolver) PlaylistByID(ctx context.Context, id int) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Where("id = ?", id).Select()

	if err != nil {
		return nil, errors.New("Failed to query playlists")
	}

	return playlists, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
