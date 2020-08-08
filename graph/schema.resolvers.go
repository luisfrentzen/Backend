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
	"math/rand"
	"strconv"
	"strings"
	"time"

	pg "github.com/go-pg/pg/v10"
)

func (r *mutationResolver) ShuffleVideos(ctx context.Context, plid string) (*model.Playlist, error) {
	var playlist model.Playlist

	err := r.DB.Model(&playlist).Where("id = ?", plid).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	s := strings.Split(playlist.Videos, ",")

	if len(s) < 2 {
		return &playlist, nil
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })

	playlist.Videos = strings.Join(s, ",")

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", plid).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &playlist, nil
}

func (r *mutationResolver) CreateLink(ctx context.Context, input *model.NewLink) (*model.Link, error) {
	link := model.Link{
		Userid: input.Userid,
		Label:  input.Label,
		URL:    input.URL,
	}

	log.Println("Inserting User")

	_, err := r.DB.Model(&link).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert new user failed")
	} else {
		log.Println("Create User Succeed")
	}

	return &link, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*model.Post, error) {
	post := model.Post{
		Userid:     input.Userid,
		Like:       input.Like,
		Disilike:   input.Disilike,
		Desc:       input.Desc,
		Attachment: input.Attachment,
		Day:        input.Day,
		Month:      input.Month,
		Year:       input.Year,
	}

	log.Println("Inserting User")

	_, err := r.DB.Model(&post).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert new user failed")
	} else {
		log.Println("Create User Succeed")
	}

	return &post, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	user := model.User{
		ID:         input.ID,
		Name:       input.Name,
		Profilepic: input.Profilepic,
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
		Premium:     input.Premium,
		Duration:    input.Duration,
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

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		Videoid:    input.Videoid,
		Userid:     input.Userid,
		Replyto:    input.Replyto,
		Like:       input.Like,
		Disilike:   input.Disilike,
		Desc:       input.Desc,
		Day:        input.Day,
		Month:      input.Month,
		Year:       input.Year,
		Replycount: input.Replycount,
		Postid:     input.Postid,
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert new comment failed")
	} else {
		log.Println("Create Comment Succeed")
	}

	if input.Videoid == 0 {
		var rootComment model.Comment

		log.Println("Getting comment")

		err := r.DB.Model(&rootComment).Where("id = ?", input.Replyto).First()

		if err != nil {
			log.Println(err)
			return nil, errors.New("Failed to query video")
		} else {
			log.Println("Query Succeed")
		}

		log.Println("Adding reply count")

		rootComment.Replycount = rootComment.Replycount + 1

		log.Println("Updating comment")

		_, updateErr := r.DB.Model(&rootComment).Where("id = ?", input.Replyto).Update()

		if updateErr != nil {
			log.Println(updateErr)
			return nil, errors.New("Update failed")
		} else {
			log.Printf("Update succeed")
		}
	}

	return &comment, nil
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

	log.Println("Updating")

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Playlist not found!")
	}

	log.Println("Playlist Found")

	if playlist.Videos == "" {
		playlist.Videos = input.Videos
	} else {
		playlist.Videos += "," + input.Videos
	}

	day, month, year := time.Now().Date()

	playlist.Day = day
	playlist.Month = int(month)
	playlist.Year = year

	log.Println("Updated Videos")

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	log.Println("Update succeeded")

	return &playlist, nil
}

func (r *mutationResolver) RemoveFromPlaylist(ctx context.Context, id string, videoid int) (*model.Playlist, error) {
	var playlist model.Playlist

	log.Println("Updating")

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Playlist not found!")
	}

	log.Println("Playlist Found")

	s := strings.Split(playlist.Videos, ",")

	for idx, i := range s {
		var j, _ = strconv.Atoi(i)
		if j == videoid {
			//log.Println(len(s))
			if len(s) == 1 {
				s = nil
			} else {
				s = append(s[:idx], s[idx+1:]...)
			}

			break
		}
	}

	day, month, year := time.Now().Date()

	playlist.Videos = strings.Join(s, ",")
	playlist.Day = day
	playlist.Month = int(month)
	playlist.Year = year

	log.Println("Updated Videos")

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	log.Println("Update succeeded")

	return &playlist, nil
}

func (r *mutationResolver) RemoveAllFromPlaylist(ctx context.Context, id string) (*model.Playlist, error) {
	var playlist model.Playlist

	log.Println("Updating")

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Playlist not found!")
	}

	log.Println("Playlist Found")

	playlist.Videos = ""

	day, month, year := time.Now().Date()

	playlist.Day = day
	playlist.Month = int(month)
	playlist.Year = year

	log.Println("Updated Videos")

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	log.Println("Update succeeded")

	return &playlist, nil
}

func (r *mutationResolver) EditPlaylist(ctx context.Context, id string, title string, visibility string, desc string) (*model.Playlist, error) {
	var playlist model.Playlist

	log.Println("Updating")

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Playlist not found!")
	}

	log.Println("Playlist Found")

	playlist.Title = title
	playlist.Visibility = visibility
	playlist.Desc = desc

	day, month, year := time.Now().Date()

	playlist.Day = day
	playlist.Month = int(month)
	playlist.Year = year

	log.Println("Updated Videos")

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	log.Println("Update succeeded")

	return &playlist, nil
}

func (r *mutationResolver) ViewVideo(ctx context.Context, id string) (*model.Video, error) {
	var video model.Video

	log.Println("Updating")

	err := r.DB.Model(&video).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Playlist not found!")
	}

	log.Println("Playlist Found")

	video.View = video.View + 1

	log.Println("Updated Videos")

	_, updateErr := r.DB.Model(&video).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	log.Println("Update succeeded")

	return &video, nil
}

func (r *mutationResolver) ViewPlaylist(ctx context.Context, id string) (*model.Playlist, error) {
	var playlist model.Playlist

	log.Println("Updating")

	err := r.DB.Model(&playlist).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Playlist not found!")
	}

	log.Println("Playlist Found")

	playlist.View = playlist.View + 1

	log.Println("Updated Videos")

	_, updateErr := r.DB.Model(&playlist).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update playlist failed")
	}

	log.Println("Update succeeded")

	return &playlist, nil
}

func (r *mutationResolver) MakeUserPremium(ctx context.Context, id string, subType string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	user.Premium = "yes"
	user.Premitype = subType
	year, month, day := time.Now().Date()
	user.Premiday = day
	user.Premimonth = int(month)
	user.Premiyear = year

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &user, nil
}

func (r *mutationResolver) Subscribe(ctx context.Context, id string, chnid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	var channel model.User

	log.Println("Getting Channel")

	err2 := r.DB.Model(&channel).Where("id = ?", chnid).First()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Channel not found!")
	}

	s := strings.Split(user.Subscribed, ",")

	if s[0] == "" {
		user.Subscribed = chnid

		channel.Subscribers = channel.Subscribers + 1
	} else {
		var subscribed = false

		for idx, i := range s {
			if i == chnid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				channel.Subscribers = channel.Subscribers - 1
				subscribed = true
				break
			}
		}

		if subscribed == false {
			s = append(s, chnid)
			channel.Subscribers = channel.Subscribers + 1
		}

		user.Subscribed = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	_, updateErr2 := r.DB.Model(&channel).Where("id = ?", chnid).Update()

	if updateErr2 != nil {
		log.Println(updateErr2)
		return nil, errors.New("Update channel failed")
	}

	return &user, nil
}

func (r *mutationResolver) Likevid(ctx context.Context, id string, chnid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	var video model.Video

	log.Println("Getting Channel")

	var vidid int
	vidid, _ = strconv.Atoi(chnid)
	err2 := r.DB.Model(&video).Where("id = ?", vidid).First()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Video not found!")
	}

	s := strings.Split(user.Likedvideos, ",")

	if s[0] == "" {
		user.Likedvideos = chnid

		video.Like = video.Like + 1
	} else {
		var liked = false

		for idx, i := range s {
			if i == chnid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				video.Like = video.Like - 1
				liked = true
				break
			}
		}

		if liked == false {
			s = append(s, chnid)
			video.Like = video.Like + 1
		}

		user.Likedvideos = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	_, updateErr2 := r.DB.Model(&video).Where("id = ?", chnid).Update()

	if updateErr2 != nil {
		log.Println(updateErr2)
		return nil, errors.New("Update channel failed")
	}

	return &user, nil
}

func (r *mutationResolver) Disilikevid(ctx context.Context, id string, chnid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	var video model.Video

	log.Println("Getting Channel")

	var vidid int
	vidid, _ = strconv.Atoi(chnid)
	err2 := r.DB.Model(&video).Where("id = ?", vidid).First()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Video not found!")
	}

	s := strings.Split(user.Disilikedvideos, ",")

	if s[0] == "" {
		user.Disilikedvideos = chnid

		video.Disilike = video.Disilike + 1
	} else {
		var liked = false

		for idx, i := range s {
			if i == chnid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				video.Disilike = video.Disilike - 1
				liked = true
				break
			}
		}

		if liked == false {
			s = append(s, chnid)
			video.Disilike = video.Disilike + 1
		}

		user.Disilikedvideos = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	_, updateErr2 := r.DB.Model(&video).Where("id = ?", chnid).Update()

	if updateErr2 != nil {
		log.Println(updateErr2)
		return nil, errors.New("Update channel failed")
	}

	return &user, nil
}

func (r *mutationResolver) Likecom(ctx context.Context, id string, chnid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	var comment model.Comment

	log.Println("Getting Comment")

	var comid int
	comid, _ = strconv.Atoi(chnid)
	err2 := r.DB.Model(&comment).Where("id = ?", comid).First()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Comment not found!")
	}

	s := strings.Split(user.Likedcomments, ",")

	if s[0] == "" {
		user.Likedcomments = chnid

		comment.Like = comment.Like + 1
	} else {
		var liked = false

		for idx, i := range s {
			if i == chnid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				comment.Like = comment.Like - 1
				liked = true
				break
			}
		}

		if liked == false {
			s = append(s, chnid)
			comment.Like = comment.Like + 1
		}

		user.Likedcomments = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	_, updateErr2 := r.DB.Model(&comment).Where("id = ?", chnid).Update()

	if updateErr2 != nil {
		log.Println(updateErr2)
		return nil, errors.New("Update comment failed")
	}

	return &user, nil
}

func (r *mutationResolver) Disilikecom(ctx context.Context, id string, chnid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	var comment model.Comment

	log.Println("Getting Comment")

	var comid int
	comid, _ = strconv.Atoi(chnid)
	err2 := r.DB.Model(&comment).Where("id = ?", comid).First()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Comment not found!")
	}

	s := strings.Split(user.Disilikedcomments, ",")

	if s[0] == "" {
		user.Disilikedcomments = chnid

		comment.Disilike = comment.Disilike + 1
	} else {
		var liked = false

		for idx, i := range s {
			if i == chnid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				comment.Disilike = comment.Disilike - 1
				liked = true
				break
			}
		}

		if liked == false {
			s = append(s, chnid)
			comment.Disilike = comment.Disilike + 1
		}

		user.Disilikedcomments = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	_, updateErr2 := r.DB.Model(&comment).Where("id = ?", chnid).Update()

	if updateErr2 != nil {
		log.Println(updateErr2)
		return nil, errors.New("Update comment failed")
	}

	return &user, nil
}

func (r *mutationResolver) Likepost(ctx context.Context, id string, chnid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	var post model.Post

	log.Println("Getting Comment")

	var curpostid int
	curpostid, _ = strconv.Atoi(chnid)
	err2 := r.DB.Model(&post).Where("id = ?", curpostid).First()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Comment not found!")
	}

	s := strings.Split(user.Likedpost, ",")

	if s[0] == "" {
		user.Likedpost = chnid

		post.Like = post.Like + 1
	} else {
		var liked = false

		for idx, i := range s {
			if i == chnid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				post.Like = post.Like - 1
				liked = true
				break
			}
		}

		if liked == false {
			s = append(s, chnid)
			post.Like = post.Like + 1
		}

		user.Likedpost = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	_, updateErr2 := r.DB.Model(&post).Where("id = ?", chnid).Update()

	if updateErr2 != nil {
		log.Println(updateErr2)
		return nil, errors.New("Update comment failed")
	}

	return &user, nil
}

func (r *mutationResolver) Disilikepost(ctx context.Context, id string, chnid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	var post model.Post

	log.Println("Getting Comment")

	var curpostid int
	curpostid, _ = strconv.Atoi(chnid)
	err2 := r.DB.Model(&post).Where("id = ?", curpostid).First()

	if err2 != nil {
		log.Println(err2)
		return nil, errors.New("Comment not found!")
	}

	s := strings.Split(user.Disilikedpost, ",")

	if s[0] == "" {
		user.Disilikedpost = chnid

		post.Disilike = post.Disilike + 1
	} else {
		var liked = false

		for idx, i := range s {
			if i == chnid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				post.Disilike = post.Disilike - 1
				liked = true
				break
			}
		}

		if liked == false {
			s = append(s, chnid)
			post.Disilike = post.Disilike + 1
		}

		user.Disilikedpost = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	_, updateErr2 := r.DB.Model(&post).Where("id = ?", chnid).Update()

	if updateErr2 != nil {
		log.Println(updateErr2)
		return nil, errors.New("Update comment failed")
	}

	return &user, nil
}

func (r *mutationResolver) Archiveplaylist(ctx context.Context, id string, plid string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	s := strings.Split(user.Archivedplaylists, ",")

	if s[0] == "" {
		user.Archivedplaylists = plid

	} else {
		var liked = false

		for idx, i := range s {
			if i == plid {
				//log.Println(len(s))
				if len(s) == 1 {
					s = nil
				} else {
					s = append(s[:idx], s[idx+1:]...)
				}

				liked = true
				break
			}
		}

		if liked == false {
			s = append(s, plid)
		}

		user.Archivedplaylists = strings.Join(s, ",")

	}

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &user, nil
}

func (r *mutationResolver) Removearchivedplaylist(ctx context.Context, id string, plid string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditAbout(ctx context.Context, id string, about string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	user.About = about

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &user, nil
}

func (r *mutationResolver) UpdateLink(ctx context.Context, id int, label string, url string) (*model.Link, error) {
	var link model.Link

	log.Println("Getting User")

	err := r.DB.Model(&link).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	link.Label = label
	link.URL = url

	_, updateErr := r.DB.Model(&link).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &link, nil
}

func (r *mutationResolver) DeleteLink(ctx context.Context, id int) (bool, error) {
	var link model.Link

	err := r.DB.Model(&link).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("User not found!")
	}

	_, deleteErr := r.DB.Model(&link).Where("id = ?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete user failed")
	}

	return true, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id int) (bool, error) {
	var post model.Post

	err := r.DB.Model(&post).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("User not found!")
	}

	_, deleteErr := r.DB.Model(&post).Where("id = ?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete user failed")
	}

	return true, nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, id int, desc string) (*model.Post, error) {
	var post model.Post

	log.Println("Getting User")

	err := r.DB.Model(&post).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	post.Desc = desc

	_, updateErr := r.DB.Model(&post).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &post, nil
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, id int) (*model.Playlist, error) {
	var pl model.Playlist

	err := r.DB.Model(&pl).Where("id = ?", id).First()

	if err != nil {
		return false, errors.New("User not found!")
	}

	_, deleteErr := r.DB.Model(&pl).Where("id = ?", id).Delete()

	if deleteErr != nil {
		return false, errors.New("Delete user failed")
	}

	return true, nil
}

func (r *mutationResolver) UpdatePlaylistSort(ctx context.Context, id int, videos string) (*model.Playlist, error) {
	var pl model.Playlist

	log.Println("Getting User")

	err := r.DB.Model(&pl).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	pl.Videos = videos

	_, updateErr := r.DB.Model(&pl).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &pl, nil
}

func (r *mutationResolver) Updateprofilepic(ctx context.Context, id string, profilepic string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	user.Profilepic = profilepic

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &user, nil
}

func (r *mutationResolver) Updatechannelart(ctx context.Context, id string, channelart string) (*model.User, error) {
	var user model.User

	log.Println("Getting User")

	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("User not found!")
	}

	user.Channelart = channelart

	_, updateErr := r.DB.Model(&user).Where("id = ?", id).Update()

	if updateErr != nil {
		log.Println(updateErr)
		return nil, errors.New("Update user failed")
	}

	return &user, nil
}

func (r *queryResolver) PostByUser(ctx context.Context, userid string) ([]*model.Post, error) {
	var posts []*model.Post

	err := r.DB.Model(&posts).Order("year DESC", "month DESC", "day DESC").Where("userid = ?", userid).Select()

	if err != nil {
		return nil, errors.New("Failed to query playlists")
	}

	return posts, nil
}

func (r *queryResolver) PostByID(ctx context.Context, id int) (*model.Post, error) {
	var post model.Post

	err := r.DB.Model(&post).Where("id = ?", id).First()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query user")
	} else {
		log.Println("Get User By Id Succeed")
	}

	return &post, nil
}

func (r *queryResolver) CommentByPost(ctx context.Context, id int) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Order("year DESC", "month DESC", "day DESC").Where("postid = ?", id).Select()

	if err != nil {
		return nil, errors.New("Failed to query playlists")
	}

	return comments, nil
}

func (r *queryResolver) LinkByUser(ctx context.Context, userid string) ([]*model.Link, error) {
	var links []*model.Link

	err := r.DB.Model(&links).Where("userid = ?", userid).Select()

	if err != nil {
		return nil, errors.New("Failed to query playlists")
	}

	return links, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Model(&users).Order("id").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query users")
	} else {
		log.Println("Get Users Succeed")
	}

	return users, nil
}

func (r *queryResolver) Videos(ctx context.Context, sort string, filter string, premium string) ([]*model.Video, error) {
	var videos []*model.Video

	if premium == "yes" {
		if sort == "view" {
			if filter == "all" {
				err := r.DB.Model(&videos).Order("view DESC").Where("visibility = ?", "public").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			} else {
				err := r.DB.Model(&videos).Order("view DESC").Where("restriction = ? and visibility = ? ", "all", "public").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			}
		} else {
			if filter == "all" {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("visibility = ?", "public").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			} else {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("restriction = ? and visibility = ?", "all", "public").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			}

		}
	} else {
		if sort == "view" {
			if filter == "all" {
				err := r.DB.Model(&videos).Order("view DESC").Where("premium = ? and visibility = ?", "no", "public").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			} else {
				err := r.DB.Model(&videos).Order("view DESC").Where("restriction = ? and visibility = ? and premium = ?", "all", "public", "no").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			}
		} else {
			if filter == "all" {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("premium = ? and visibility = ?", "no", "public").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			} else {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("restriction = ? and visibility = ? and premium = ?", "all", "public", "no").Select()

				if err != nil {
					log.Println(err)
					return nil, errors.New("Failed to query videos")
				}
			}

		}
	}

	return videos, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Order("year DESC", "month DESC", "day DESC").Select()

	if err != nil {
		return nil, errors.New("Failed to query videos")
	}

	return comments, nil
}

func (r *queryResolver) CommentsByVideo(ctx context.Context, videoid int, sort string) ([]*model.Comment, error) {
	var comments []*model.Comment

	if sort == "" {
		err := r.DB.Model(&comments).Order("year DESC", "month DESC", "day DESC").Where("videoid = ?", videoid).Select()

		if err != nil {
			log.Println(err)
			return nil, errors.New("Failed to query playlists")
		}
	} else if sort == "like" {
		err := r.DB.Model(&comments).Order("like DESC").Where("videoid = ?", videoid).Select()

		if err != nil {
			log.Println(err)
			return nil, errors.New("Failed to query playlists")
		}
	}

	return comments, nil
}

func (r *queryResolver) Replies(ctx context.Context, replyto int) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Order("year DESC", "month DESC", "day DESC").Where("replyto = ?", replyto).Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query playlists")
	}

	return comments, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Order("year DESC", "month DESC", "day DESC").Order("id").Select()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query videos")
	}

	return playlists, nil
}

func (r *queryResolver) PlaylistsByUser(ctx context.Context, userid string, visibility string) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Order("visibility ASC").Where("userid = ? and visibility != ?", userid, visibility).Select()

	if err != nil {
		return nil, errors.New("Failed to query playlists")
	}

	return playlists, nil
}

func (r *queryResolver) VideosByUser(ctx context.Context, userid string, sort string, premium string, privacy string) ([]*model.Video, error) {
	var videos []*model.Video

	if premium == "yes" {
		if privacy == "all" {
			if sort == "" {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("userid = ?", userid).Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "old" {
				err := r.DB.Model(&videos).Order("year ASC", "month ASC", "day ASC").Where("userid = ?", userid).Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "view" {
				err := r.DB.Model(&videos).Order("view DESC").Where("userid = ? ", userid).Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			}
		} else {
			if sort == "" {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("userid = ? and visibility = ?", userid, "public").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "old" {
				err := r.DB.Model(&videos).Order("year ASC", "month ASC", "day ASC").Where("userid = ? and visibility = ?", userid, "public").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "view" {
				err := r.DB.Model(&videos).Order("view DESC").Where("userid = ? and visibility = ?", userid, "public").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			}
		}
	} else {
		if privacy == "all" {
			if sort == "" {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("userid = ? and premium = ?", userid, "no").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "old" {
				err := r.DB.Model(&videos).Order("year ASC", "month ASC", "day ASC").Where("userid = ? and premium = ? ", userid, "no").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "view" {
				err := r.DB.Model(&videos).Order("view DESC").Where("userid = ? and premium = ? ", userid, "no").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			}
		} else {
			if sort == "" {
				err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("userid = ? and premium = ? and visibility = ?", userid, "no", "public").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "old" {
				err := r.DB.Model(&videos).Order("year ASC", "month ASC", "day ASC").Where("userid = ? and premium = ? and visibility = ?", userid, "no", "public").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			} else if sort == "view" {
				err := r.DB.Model(&videos).Order("view DESC").Where("userid = ? and premium = ? and visibility = ?", userid, "no", "public").Select()

				if err != nil {
					return nil, errors.New("Failed to query videos")
				}
			}
		}
	}

	return videos, nil
}

func (r *queryResolver) VideosByCategory(ctx context.Context, category string) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Order("year DESC", "month DESC", "day DESC").Where("category = ?", category).Select()

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

func (r *queryResolver) VideosByIds(ctx context.Context, id string) ([]*model.Video, error) {
	var videos []*model.Video

	s := strings.Split(id, ",")
	var idArr = []int64{}

	for _, i := range s {
		j, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			log.Println(err)
		}

		idArr = append(idArr, j)
	}

	_, err := r.DB.Query(&videos, `SELECT * FROM videos WHERE id IN (?) ORDER BY videos.view ASC`, pg.Ints(idArr))

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query video")
	} else {
		log.Println("Get Video By Id Succeed")
	}
	return videos, nil
}

func (r *queryResolver) CommentByID(ctx context.Context, id int) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Where("id = ?", id).Select()

	if err != nil {
		return nil, errors.New("Failed to query comment")
	}

	return comments, nil
}

func (r *queryResolver) VideosByUsers(ctx context.Context, id string, premium string) ([]*model.Video, error) {
	var videos []*model.Video

	var s []string
	s = strings.Split(id, ",")

	if premium == "yes" {
		_, err := r.DB.Query(&videos, `SELECT * FROM videos WHERE userid IN (?) AND videos.visibility = (?) ORDER BY videos.year DESC`, pg.Strings(s), "public")
		if err != nil {
			log.Println(err)
			return nil, errors.New("Failed to query videos")
		}
	} else {
		_, err := r.DB.Query(&videos, `SELECT * FROM videos WHERE userid IN (?) AND videos.visibility = (?) AND videos.premium = (?) ORDER BY videos.year DESC, videos.month DESC, videos.day DESC`, pg.Strings(s), "public", "no")
		if err != nil {
			log.Println(err)
			return nil, errors.New("Failed to query videos")
		}
	}

	return videos, nil
}

func (r *queryResolver) UsersByIds(ctx context.Context, id string) ([]*model.User, error) {
	var users []*model.User

	s := strings.Split(id, ",")
	//var idArr = []int64{}

	//for _, i := range s {
	//j, err := strconv.ParseInt(i, 10, 64)
	//if err != nil {
	//	log.Println(err)
	//}

	//	idArr = append(idArr, j)
	//}

	_, err := r.DB.Query(&users, `SELECT * FROM users WHERE id IN (?) ORDER BY users.name ASC`, pg.Strings(s))

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query video")
	} else {
		log.Println("Get Video By Id Succeed")
	}
	return users, nil
}

func (r *queryResolver) GetArchivedPlaylist(ctx context.Context, ids string) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	s := strings.Split(ids, ",")
	var idArr = []int64{}

	for _, i := range s {
		j, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			log.Println(err)
		}

		idArr = append(idArr, j)
	}

	_, err := r.DB.Query(&playlists, `SELECT * FROM playlists WHERE id IN (?)`, pg.Ints(idArr))

	if err != nil {
		log.Println(err)
		return nil, errors.New("Failed to query video")
	} else {
		log.Println("Get Video By Id Succeed")
	}
	return playlists, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
