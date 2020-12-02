package newsfeed

import "testing"

func TestAdd(t *testing.T){
	feed := New()
	feed.Add(Item{})
	if len(feed.Items) != 1{
		t.Errorf("No Items")
	}

} 

func TestGetAll(t *testing.T){
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1{
		t.Errorf("No Items")
	}

}